package cashflowout

import (
	"time"

	"github.com/google/uuid"
)

type CashFlowBalance struct {
	ID           uuid.UUID `json:"id"`
	Date         time.Time `json:"date"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	ComissionID  *string   `json:"comission_id"`
	DueDate      time.Time `json:"due_date"`
	PaidDate     time.Time `json:"paid_date"`
	Paid         bool      `json:"paid"`
	Value        float64   `json:"value"`
	Type         string    `json:"type"`
	Description  string    `json:"description"`
}
type CashFlowBalanceParams struct {
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	ComissionID  *string   `json:"comission_id"`
	DueDate      time.Time `json:"due_date"`
	PaidDate     time.Time `json:"paid_date"`
	Paid         bool      `json:"paid"`
	Value        float64   `json:"value"`
	Type         string    `json:"type"`
	Description  string    `json:"description"`
}
type CashFlowBalanceUpdate struct {
	ID           uuid.UUID `json:"id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	ComissionID  *string   `json:"comission_id"`
	DueDate      *time.Time `json:"due_date"`
	PaidDate     *time.Time `json:"paid_date"`
	Paid         *bool      `json:"paid"`
	Value        *float64   `json:"value"`
	Type         *string    `json:"type"`
	Description  *string    `json:"description"`
}

func NewCashFlowBalance(p *CashFlowBalanceParams) *CashFlowBalance {
	return &CashFlowBalance{
		ID: uuid.New(),
		Date: time.Now(),
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		ComissionID: p.ComissionID,
		DueDate: p.DueDate,
		PaidDate: p.PaidDate,
		Paid: p.Paid,
		Value: p.Value,
		Type: p.Type,
		Description: p.Description,
	}
}