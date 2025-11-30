package postgres

import (
	"context"

	"github.com/choiexe1/go-ecom/internal/products"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
)

type repository struct {
	queries *repo.Queries
}

func NewRepository(queries *repo.Queries) products.Repository {
	return &repository{queries: queries}
}

func (r *repository) FindAll(ctx context.Context) ([]products.Product, error) {
	rows, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]products.Product, len(rows))
	for i, row := range rows {
		result[i] = toProduct(row)
	}
	return result, nil
}

func (r *repository) FindByID(ctx context.Context, id int64) (products.Product, error) {
	row, err := r.queries.GetProductByID(ctx, id)
	if err != nil {
		return products.Product{}, err
	}
	return toProduct(row), nil
}

func (r *repository) Create(ctx context.Context, params products.CreateProductParams) (products.Product, error) {
	row, err := r.queries.CreateProduct(ctx, repo.CreateProductParams{
		Name:       params.Name,
		PriceCents: params.PriceCents,
		Quantity:   params.Quantity,
	})
	if err != nil {
		return products.Product{}, err
	}
	return toProduct(row), nil
}

func (r *repository) Update(ctx context.Context, params products.UpdateProductParams) (products.Product, error) {
	row, err := r.queries.UpdateProduct(ctx, repo.UpdateProductParams{
		ID:         params.ID,
		Name:       params.Name,
		PriceCents: params.PriceCents,
		Quantity:   params.Quantity,
	})
	if err != nil {
		return products.Product{}, err
	}
	return toProduct(row), nil
}

func toProduct(p repo.Product) products.Product {
	return products.Product{
		ID:         p.ID,
		Name:       p.Name,
		PriceCents: p.PriceCents,
		Quantity:   p.Quantity,
		CreatedAt:  p.CreatedAt.Time,
	}
}
