package products

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id int64) (Product, error)
	Create(ctx context.Context, params CreateProductParams) (Product, error)
	Update(ctx context.Context, params UpdateProductParams) (Product, error)
}
