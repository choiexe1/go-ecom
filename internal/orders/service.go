package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (PlaceOrderResponse, error) {
	if tempOrder.CustomerID == 0 {
		return PlaceOrderResponse{}, fmt.Errorf("customer ID is required")
	}
	if len(tempOrder.Items) == 0 {
		return PlaceOrderResponse{}, fmt.Errorf("at least one item is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return PlaceOrderResponse{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return PlaceOrderResponse{}, err
	}

	var orderItems []repo.OrderItem
	var totalPrice int32

	for _, item := range tempOrder.Items {
		product, err := qtx.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return PlaceOrderResponse{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return PlaceOrderResponse{}, ErrProductNoStock
		}

		product.Quantity = product.Quantity - item.Quantity

		if product.Quantity < 0 {
			product.Quantity = 0
		}

		orderItem, err := qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceCents,
		})
		if err != nil {
			return PlaceOrderResponse{}, err
		}
		orderItems = append(orderItems, orderItem)
		totalPrice += orderItem.PriceCents * orderItem.Quantity

		qtx.UpdateProduct(ctx, repo.UpdateProductParams{
			ID:         item.ProductID,
			Name:       product.Name,
			PriceCents: product.PriceCents,
			Quantity:   product.Quantity,
		})
	}

	tx.Commit(ctx)

	return PlaceOrderResponse{
		Order:      order,
		Items:      orderItems,
		TotalPrice: totalPrice,
	}, nil
}
