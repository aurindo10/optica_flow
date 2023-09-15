package comission

import (
	"time"

	"github.com/google/uuid"
)



type Commission struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	OrderID      *string   `json:"order_id"`
	Value        float64   `json:"value"`
}


type CommissionParams struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	OrderID      *string   `json:"order_id"`
	Value        float64   `json:"value"`
}
type CommissionToUpdate struct {
	ID           uuid.UUID `json:"id"`
	Name         *string    `json:"name"`
	Description  *string    `json:"description"`
	WhoUpdatedID *string    `json:"who_updated_id"`
	Value        *float64   `json:"value"`
	UpdatedAt    time.Time `json:"updated_at"`
}


func NewComission(p *CommissionParams) *Commission {
	return &Commission{
		ID: uuid.New(),
		Name: p.Name,
		Description: p.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OrderID: p.OrderID,
		Value: p.Value,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
	}
}