package flowentriesrepository

import (
	"context"
	flowentries "optica_flow/internal/app/domain/flow_entries"
	database "optica_flow/internal/app/infra/database/queries"
)

type FlowEntriesRepository struct {
	db *database.Queries
}

func (c *FlowEntriesRepository) Create(p *flowentries.CashFlowEntries) (*flowentries.CashFlowEntries, error) {
	request := database.CreateCashFlowEntrieParams{
		Type: p.Type,
		Amount: p.Amount,
		Description: p.Description,
		CompanyID: p.CompanyID,
		OrderID: p.OrderID,
		WhoCreatedID: p.WhoCreatedID,	
		ID: p.ID,
	}
	result, error := c.db.CreateCashFlowEntrie(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &flowentries.CashFlowEntries{
		ID: result.ID,
		Date: result.Date,
		Type: result.Type,
		Amount: result.Amount,
		Description: result.Description,
		CompanyID: result.CompanyID,
		OrderID: result.OrderID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
	}
	return response, nil
}
func (c *FlowEntriesRepository) FindByIntervalDate(p *flowentries.FindByDateParams) ([]*flowentries.CashFlowEntries, error){
	request := database.GetCashFlowEntriesByDateRangeAndCompanyParams{
		CompanyID: p.CompanyId,
		Date: p.InitialDate,
		Date_2: p.EndDate,
	}
	result, error := c.db.GetCashFlowEntriesByDateRangeAndCompany(context.Background(), request)
	if error != nil {
		return nil, error
	}
	var response []*flowentries.CashFlowEntries
	for _, flowEntry := range result {
		response = append(response, &flowentries.CashFlowEntries{
			ID: flowEntry.ID,
			Date: flowEntry.Date,
			Type: flowEntry.Type,
			Amount: flowEntry.Amount,
			Description: flowEntry.Description,
			CompanyID: flowEntry.CompanyID,
			OrderID: flowEntry.OrderID,
			WhoCreatedID: flowEntry.WhoCreatedID,
			WhoUpdatedID: flowEntry.WhoUpdatedID,
		})
	}
	return response, nil
}
func (c *FlowEntriesRepository) Update(p *flowentries.CashFlowEntriesUpdate) (*flowentries.CashFlowEntries, error){
	request := database.UpdateFlowEntrieParams{
		ID: p.ID,
		Type: p.Type,
		Amount: p.Amount,
		Description: p.Description,
		WhoUpdatedID: p.WhoUpdatedID,
	}
	result, error := c.db.UpdateFlowEntrie(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &flowentries.CashFlowEntries{
		ID: result.ID,
		Date: result.Date,
		Type: result.Type,
		Amount: result.Amount,
		Description: result.Description,
		CompanyID: result.CompanyID,
		OrderID: result.OrderID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
	}
	return response, nil 
}
func NewFlowEntrieRepository(db *database.Queries) *FlowEntriesRepository {
	return &FlowEntriesRepository{
		db : db,
	}
}