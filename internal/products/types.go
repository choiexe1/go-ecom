package products

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	PriceCents int32     `json:"priceCents"`
	Quantity   int32     `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
}

type UpdateProductParams struct {
	ID         int64
	Name       string
	PriceCents int32
	Quantity   int32
}

type CreateProductParams struct {
	Name       string `json:"name"`
	PriceCents int32  `json:"priceCents"`
	Quantity   int32  `json:"quantity"`
}