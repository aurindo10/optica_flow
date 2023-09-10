package productorder

import "github.com/google/uuid"

type ProductOrder struct {
	ID        uuid.UUID `json:"id"`
	Amout     int32     `json:"amout"`
	ProductID *string   `json:"product_id"`
	OrderID   *string   `json:"order_id"`
}
type ProductOrderParms struct {
	ID        uuid.UUID `json:"id"`
	Amout     int32     `json:"amout"`
	ProductID *string   `json:"product_id"`
	OrderID   *string   `json:"order_id"`
}

func NewProductOrder(product *ProductOrderParms) *ProductOrder {
	return &ProductOrder{
		ID:        uuid.New(),
		Amout:     product.Amout,
		ProductID: product.ProductID,
		OrderID:   product.OrderID,
	}
}