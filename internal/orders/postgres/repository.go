package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
	"github.com/choiexe1/go-ecom/internal/orders"
	"github.com/choiexe1/go-ecom/internal/products"
)

type repository struct {
	queries *repo.Queries
	db      *pgx.Conn
}

func NewRepository(queries *repo.Queries, db *pgx.Conn) orders.Repository {
	return &repository{
		queries: queries,
		db:      db,
	}
}

func (r *repository) WithTx(fn func(tx orders.TxRepository) error) error {
	pgxTx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer pgxTx.Rollback(context.Background())

	txRepo := &txRepository{queries: r.queries.WithTx(pgxTx)}
	if err := fn(txRepo); err != nil {
		return err
	}

	return pgxTx.Commit(context.Background())
}

type txRepository struct {
	queries *repo.Queries
}

func (t *txRepository) CreateOrder(ctx context.Context, customerID int64) (orders.Order, error) {
	row, err := t.queries.CreateOrder(ctx, customerID)
	if err != nil {
		return orders.Order{}, err
	}
	return orders.Order{
		ID:         row.ID,
		CustomerID: row.CustomerID,
		CreatedAt:  row.CreatedAt.Time,
	}, nil
}

func (t *txRepository) CreateOrderItem(ctx context.Context, params orders.CreateOrderItemParams) (orders.OrderItem, error) {
	row, err := t.queries.CreateOrderItem(ctx, repo.CreateOrderItemParams{
		OrderID:    params.OrderID,
		ProductID:  params.ProductID,
		Quantity:   params.Quantity,
		PriceCents: params.PriceCents,
	})
	if err != nil {
		return orders.OrderItem{}, err
	}
	return orders.OrderItem{
		ID:         row.ID,
		OrderID:    row.OrderID,
		ProductID:  row.ProductID,
		Quantity:   row.Quantity,
		PriceCents: row.PriceCents,
	}, nil
}

func (t *txRepository) GetProductByID(ctx context.Context, id int64) (products.Product, error) {
	row, err := t.queries.GetProductByID(ctx, id)
	if err != nil {
		return products.Product{}, err
	}
	return products.Product{
		ID:         row.ID,
		Name:       row.Name,
		PriceCents: row.PriceCents,
		Quantity:   row.Quantity,
		CreatedAt:  row.CreatedAt.Time,
	}, nil
}

func (t *txRepository) UpdateProduct(ctx context.Context, params products.UpdateProductParams) (products.Product, error) {
	row, err := t.queries.UpdateProduct(ctx, repo.UpdateProductParams{
		ID:         params.ID,
		Name:       params.Name,
		PriceCents: params.PriceCents,
		Quantity:   params.Quantity,
	})
	if err != nil {
		return products.Product{}, err
	}
	return products.Product{
		ID:         row.ID,
		Name:       row.Name,
		PriceCents: row.PriceCents,
		Quantity:   row.Quantity,
		CreatedAt:  row.CreatedAt.Time,
	}, nil
}
