package flowentries

import (
	"time"

	"github.com/google/uuid"
)

type CashFlowEntries struct {
	ID           uuid.UUID `json:"id"`
	Date         time.Time `json:"date"`
	Type         string    `json:"type"`
	Amount       float64   `json:"amount"`
	Description  string    `json:"description"`
	CompanyID    string    `json:"company_id"`
	OrderID      *string   `json:"order_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}
type CashFlowEntriesParams struct {
	ID           uuid.UUID `json:"id"`
	Date         time.Time `json:"date"`
	Type         string    `json:"type"`
	Amount       float64   `json:"amount"`
	Description  string    `json:"description"`
	CompanyID    string    `json:"company_id"`
	OrderID      *string   `json:"order_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}

func NewCashFlowEntries(p *CashFlowEntriesParams) *CashFlowEntries {
	return &CashFlowEntries{
		ID: uuid.New(),
		Date: p.Date,
		Type: p.Type,
		Amount: p.Amount,
		Description: p.Description,
		CompanyID: p.CompanyID,
		OrderID: p.OrderID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
	}
}