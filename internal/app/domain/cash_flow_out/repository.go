package cashflowout

import (
	"time"

	"github.com/google/uuid"
)


type FindByRangeDateParams struct {
	InitialDate time.Time 	`json:"date"`
	EndDate time.Time		`json:"date_2"`
	CompanyId string		`json:"company_id"`
}
type Repository interface {
	Create(*CashFlowBalance) (*CashFlowBalance, error)
	FindByRangeDate(*FindByRangeDateParams) ([]*CashFlowBalance, error)
	Update(*CashFlowBalanceUpdate) (*CashFlowBalance, error)
	Delete(id uuid.UUID) error
}