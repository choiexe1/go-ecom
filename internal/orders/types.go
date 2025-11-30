package orders

import (
	"context"
	"time"

	"github.com/choiexe1/go-ecom/internal/products"
)

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customerId"`
	CreatedAt  time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID         int64 `json:"id"`
	OrderID    int64 `json:"orderId"`
	ProductID  int64 `json:"productId"`
	Quantity   int32 `json:"quantity"`
	PriceCents int32 `json:"priceCents"`
}

type CreateOrderItemParams struct {
	OrderID    int64
	ProductID  int64
	Quantity   int32
	PriceCents int32
}

type OrderItemRequest struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type CreateOrderParams struct {
	CustomerID int64              `json:"customerId"`
	Items      []OrderItemRequest `json:"items"`
}

type PlaceOrderResponse struct {
	Order      Order       `json:"order"`
	Items      []OrderItem `json:"items"`
	TotalPrice int32       `json:"totalPrice"`
}

type Service interface {
	PlaceOrder(ctx context.Context, params CreateOrderParams) (PlaceOrderResponse, error)
}

type Repository interface {
	WithTx(fn func(tx TxRepository) error) error
}

type TxRepository interface {
	CreateOrder(ctx context.Context, customerID int64) (Order, error)
	CreateOrderItem(ctx context.Context, params CreateOrderItemParams) (OrderItem, error)
	GetProductByID(ctx context.Context, id int64) (products.Product, error)
	UpdateProduct(ctx context.Context, params products.UpdateProductParams) (products.Product, error)
}
