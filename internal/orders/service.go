package orders

import (
	"context"
	"errors"
	"fmt"

	"github.com/choiexe1/go-ecom/internal/products"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product has not enough stock")
)

type svc struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, params CreateOrderParams) (PlaceOrderResponse, error) {
	if params.CustomerID == 0 {
		return PlaceOrderResponse{}, fmt.Errorf("customer ID is required")
	}
	if len(params.Items) == 0 {
		return PlaceOrderResponse{}, fmt.Errorf("at least one item is required")
	}

	var response PlaceOrderResponse

	err := s.repo.WithTx(func(tx TxRepository) error {
		order, err := tx.CreateOrder(ctx, params.CustomerID)
		if err != nil {
			return err
		}

		var orderItems []OrderItem
		var totalPrice int32

		for _, item := range params.Items {
			product, err := tx.GetProductByID(ctx, item.ProductID)
			if err != nil {
				return ErrProductNotFound
			}

			if product.Quantity < item.Quantity {
				return ErrProductNoStock
			}

			newQuantity := product.Quantity - item.Quantity
			if newQuantity < 0 {
				newQuantity = 0
			}

			orderItem, err := tx.CreateOrderItem(ctx, CreateOrderItemParams{
				OrderID:    order.ID,
				ProductID:  item.ProductID,
				Quantity:   item.Quantity,
				PriceCents: product.PriceCents,
			})
			if err != nil {
				return err
			}

			orderItems = append(orderItems, orderItem)
			totalPrice += orderItem.PriceCents * orderItem.Quantity

			_, err = tx.UpdateProduct(ctx, products.UpdateProductParams{
				ID:         item.ProductID,
				Name:       product.Name,
				PriceCents: product.PriceCents,
				Quantity:   newQuantity,
			})
			if err != nil {
				return err
			}
		}

		response = PlaceOrderResponse{
			Order:      order,
			Items:      orderItems,
			TotalPrice: totalPrice,
		}
		return nil
	})

	if err != nil {
		return PlaceOrderResponse{}, err
	}

	return response, nil
}
