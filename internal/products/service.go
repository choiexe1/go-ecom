package products

import (
	"context"
	"errors"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *svc) GetProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)

	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}

func (s *svc) UpdateProduct(ctx context.Context, product repo.Product) (repo.Product, error) {
	product, err := s.repo.UpdateProduct(ctx, repo.UpdateProductParams{
		ID:         product.ID,
		Name:       product.Name,
		PriceCents: product.PriceCents,
		Quantity:   product.Quantity,
	})

	if err != nil {
		return repo.Product{}, err
	}

	return product, nil
}
