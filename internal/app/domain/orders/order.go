package orders

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `json:"id"`
	ProductName  string    `json:"product_name"`
	Quantity     int32     `json:"quantity"`
	OrderDate    time.Time `json:"order_date"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	ClientID     *string   `json:"client_id"`
	CompanyID    string    `json:"company_id"`
	Fase         string    `json:"fase"`
}
type Params struct {
	ProductName  string    `json:"product_name"`
	Quantity     int32     `json:"quantity"`
	OrderDate    time.Time `json:"order_date"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	ClientID     *string   `json:"client_id"`
	CompanyID    string    `json:"company_id"`
	Fase         string    `json:"fase"`
}
type OrderToUpdate struct {
	ID           uuid.UUID `json:"id"`
	ProductName  *string   `json:"product_name"`
	Quantity     *int32    `json:"quantity"`
	WhoUpdatedID *string   `json:"who_updated_id"`
	Fase         *string   `json:"fase"`
}
func NewOrder(p *Params) *Order {
	return &Order{
		ID:           uuid.New(),
		ProductName:  p.ProductName,
		Quantity:     p.Quantity,
		OrderDate:    p.OrderDate,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		ClientID:     p.ClientID,
		CompanyID:    p.CompanyID,
		Fase:         p.Fase,
	}
}