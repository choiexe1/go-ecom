package orders

import (
	"context"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (PlaceOrderResponse, error)
}

type PlaceOrderResponse struct {
	Order      repo.Order       `json:"order"`
	Items      []repo.OrderItem `json:"items"`
	TotalPrice int32            `json:"totalPrice"`
}
