package points

import (
	"time"

	"github.com/google/uuid"
)


type Points struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Ammount     float64   `json:"ammount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ValidDate   time.Time `json:"valid_date"`
	CompanyID   string    `json:"company_id"`
	OrderID     *string   `json:"order_id"`
}
type PointsParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Ammount     float64   `json:"ammount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ValidDate   time.Time `json:"valid_date"`
	CompanyID   string    `json:"company_id"`
	OrderID     string   `json:"order_id"`
}

func NewPoints(p *PointsParams) *Points {
	return &Points{
		ID: uuid.New(),
		Name: p.Name,
		Description: p.Description,
		Active: p.Active,
		Ammount: p.Ammount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ValidDate: p.ValidDate,
		CompanyID: p.CompanyID,
		OrderID: &p.OrderID,
	}
}