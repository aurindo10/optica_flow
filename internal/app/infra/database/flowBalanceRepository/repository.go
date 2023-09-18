package flowbalancerepository

import (
	"context"
	cashflowout "optica_flow/internal/app/domain/cash_flow_out"
	database "optica_flow/internal/app/infra/database/queries"
)




type FlowBalanceRepository struct {
	db *database.Queries
}
func (c *FlowBalanceRepository) Create(p *cashflowout.CashFlowBalance) (*cashflowout.CashFlowBalance, error) {
	request := database.CreateFlowBalanceParams{
		ID: p.ID,
		Date: p.Date,
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
	result, error := c.db.CreateFlowBalance(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := cashflowout.CashFlowBalance{
		ID: result.ID,
		Date: result.Date,
		CompanyID: result.CompanyID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
		ComissionID: result.ComissionID,
		DueDate: result.DueDate,
		PaidDate: result.PaidDate,
		Paid: result.Paid,
		Value: result.Value,
		Type: result.Type,
		Description: result.Description,
	}
	return &response, nil
}

func (c *FlowBalanceRepository) FindByRangeDate(p *cashflowout.FindByRangeDateParams) ([]*cashflowout.CashFlowBalance, error) {
	request := database.FindFlowBalanceByDateRangeParams{
		DueDate: p.InitialDate,
		DueDate_2: p.EndDate,
		CompanyID: p.CompanyId,
	}
	result, error := c.db.FindFlowBalanceByDateRange(context.Background(), request)
	if error != nil {
		return nil, error
	}
	var response []*cashflowout.CashFlowBalance
	for _, item := range result {
		response = append(response, &cashflowout.CashFlowBalance{
			ID: item.ID,
			Date: item.Date,
			CompanyID: item.CompanyID,
			WhoCreatedID: item.WhoCreatedID,
			WhoUpdatedID: item.WhoUpdatedID,
			ComissionID: item.ComissionID,
			DueDate: item.DueDate,
			PaidDate: item.PaidDate,
			Paid: item.Paid,
			Value: item.Value,
			Type: item.Type,
			Description: item.Description,
		} )
	}
	return response, nil
}

func NewFlowBalanceRepository(db *database.Queries) *FlowBalanceRepository {
	return &FlowBalanceRepository{
		db: db,
	}
}