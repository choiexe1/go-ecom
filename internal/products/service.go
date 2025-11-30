package products

import (
	"context"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type Service interface {
	ListProducts(ctx context.Context) ([]Product, error)
	GetProductByID(ctx context.Context, id int64) (Product, error)
	CreateProduct(ctx context.Context, params CreateProductParams) (Product, error)
}

type svc struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *svc) GetProductByID(ctx context.Context, id int64) (Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *svc) CreateProduct(ctx context.Context, params CreateProductParams) (Product, error) {
	return s.repo.Create(ctx, params)
}
