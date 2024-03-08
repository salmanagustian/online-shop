package products

import (
	"context"
	"database/sql"
	"online-shop/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			sku, name, stock , price, created_at, updated_at
		) VALUES (
			:sku, :name, :stock, :price, :created_at, :updated_at
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)

	return
}

// GetAllProductWithPaginationCursor implements Repository.
func (r repository) GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := `
		SELECT
			id, sku, name
			, stock, price
			, created_at
			, updated_at
		FROM products
		WHERE id>$1 
		ORDER BY id ASC
		LIMIT $2
	`

	err = r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}

		return
	}

	return
}

// GetProductBySku implements Repository.
func (r repository) GetProductBySku(ctx context.Context, sku string) (product Product, err error) {
	query := `
	SELECT
		id, sku, name
		, stock, price
		, created_at
		, updated_at
	FROM products
	WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

// UpdateProductById implements Repository.
func (r repository) UpdateProductById(ctx context.Context, model Product, id int) (err error) {
	var productId int
	// find product
	queryFindProduct := `SELECT id from products WHERE id=$1`

	row := r.db.QueryRowxContext(ctx, queryFindProduct, id)

	if err = row.Scan(&productId); err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	queryUpdate := `UPDATE products SET name=$1, stock=$2, price=$3 WHERE id=$4`

	_, err = r.db.ExecContext(ctx, queryUpdate, model.Name, model.Stock, model.Price, productId)

	if err != nil {
		return
	}

	return
}

// DeleteProductById implements Repository.
func (r repository) DeleteProductById(ctx context.Context, id int) (err error) {
	var productId int
	// find product
	queryFindProduct := `SELECT id from products WHERE id=$1`

	row := r.db.QueryRowxContext(ctx, queryFindProduct, id)

	if err = row.Scan(&productId); err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	queryDelete := `DELETE from products WHERE id=$1`

	_, err = r.db.ExecContext(ctx, queryDelete, productId)

	if err != nil {
		return
	}
	return
}
