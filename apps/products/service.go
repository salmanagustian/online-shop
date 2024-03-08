package products

import (
	"context"
	"online-shop/infra/response"
	"online-shop/internal/log"
)

type Repository interface {
	CreateProduct(ctx context.Context, model Product) (err error)
	GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error)
	GetProductBySku(ctx context.Context, sku string) (product Product, err error)
	UpdateProductById(ctx context.Context, model Product, id int) (err error)
	DeleteProductById(ctx context.Context, id int) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := NewProductFromCreateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		log.Log.Errorf(ctx, "[CreateProduct, Validate] with error detail: %v", err.Error())
		return
	}

	if err = s.repo.CreateProduct(ctx, productEntity); err != nil {
		return
	}

	return
}

func (s service) ListProducts(ctx context.Context, req ListProductRequestPayload) (products []Product, err error) {
	pagination := NewProductPaginationFromListProductRequest(req)

	products, err = s.repo.GetAllProductWithPaginationCursor(ctx, pagination)

	if err != nil {
		if err == response.ErrNotFound {
			return []Product{}, nil
		}
		return
	}

	if len(products) == 0 {
		return []Product{}, nil
	}

	return
}

func (s service) ProductDetail(ctx context.Context, sku string) (model Product, err error) {
	model, err = s.repo.GetProductBySku(ctx, sku)
	if err != nil {
		return
	}
	return
}

func (s service) DeleteProduct(ctx context.Context, id int) (err error) {
	err = s.repo.DeleteProductById(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s service) UpdateProduct(ctx context.Context, req UpdateProductRequestPayload, id int) (err error) {
	productEntity := NewProductFromUpdateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		return
	}

	if err = s.repo.UpdateProductById(ctx, productEntity, id); err != nil {
		return
	}
	return
}
