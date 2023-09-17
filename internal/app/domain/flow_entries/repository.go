package flowentries

import (
	"time"

	"github.com/google/uuid"
)

type FindByDateParams struct {
	InitialDate time.Time 	`json:"date"`
	EndDate time.Time		`json:"date_2"`
	CompanyId string		`json:"company_id"`
}

type Repository interface {
	Create(*CashFlowEntries) (*CashFlowEntries, error)
	FindByIntervalDate(*FindByDateParams) ([]*CashFlowEntries, error)
	Update(*CashFlowEntriesUpdate) (*CashFlowEntries, error)
	Delete(id uuid.UUID) error
}

