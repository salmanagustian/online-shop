package products

import (
	"context"
	"log"
	"online-shop/external/database"
	"online-shop/infra/response"
	"online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Succcess(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Baju baru",
		Stock: 10,
		Price: 10_000,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Stock: 10,
			Price: 10_000,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestProductDetail_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Baju baru",
		Stock: 10,
		Price: 10_000,
	}

	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})

	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)
	log.Printf("%+v", product)
}

func TestDeleteProduct_Success(t *testing.T) {
	// preparation
	req := CreateProductRequestPayload{
		Name:  "Baju baru",
		Stock: 10,
		Price: 10_000,
	}

	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})

	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)
	// end preparation

	err = svc.DeleteProduct(ctx, products[len(products)-1].Id)
	require.Nil(t, err)
	log.Printf("deleted product %v", products[len(products)-1])
}

func TestDeleteProduct_Fail(t *testing.T) {
	productId := 99999

	ctx := context.Background()

	err := svc.DeleteProduct(ctx, productId)

	require.NotNil(t, err)
	require.Equal(t, response.ErrNotFound, err)
}

func TestUpdateProduct_Success(t *testing.T) {
	// preparation
	ctx := context.Background()

	reqCreate := CreateProductRequestPayload{
		Name:  "Baju baru",
		Stock: 10,
		Price: 10_000,
	}

	err := svc.CreateProduct(ctx, reqCreate)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})

	require.Nil(t, err)
	require.NotEmpty(t, products)
	require.Greater(t, len(products), 0)
	// end preparation

	reqUpdate := UpdateProductRequestPayload{
		Name:  "baju baru updated",
		Stock: 10,
		Price: 10_000,
	}

	err = svc.UpdateProduct(ctx, reqUpdate, products[0].Id)
	require.Nil(t, err)
}

func TestUpdateProduct_Fail(t *testing.T) {
	productId := 99999

	ctx := context.Background()

	err := svc.DeleteProduct(ctx, productId)

	require.NotNil(t, err)
	require.Equal(t, response.ErrNotFound, err)
}
