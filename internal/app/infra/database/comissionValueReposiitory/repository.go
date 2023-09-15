package comissionvaluerepository

import (
	"context"
	comissionvalue "optica_flow/internal/app/domain/comission_value"
	database "optica_flow/internal/app/infra/database/queries"
)




type ComissionValueRepository struct {
	db *database.Queries
}

func (c *ComissionValueRepository) Create(p*comissionvalue.ComissionValues) (*comissionvalue.ComissionValues, error) {
	request := database.CreateComissionValueParams{
		ID: p.ID,
		Percentage: p.Percentage,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		Type: p.Type,
	}
	result, error := c.db.CreateComissionValue(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &comissionvalue.ComissionValues{
		ID: result.ID,
		Percentage: result.Percentage,
		CompanyID: result.CompanyID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
		Type: result.Type,
	}
	return response, nil
}

func NewComissionValueRepository(db *database.Queries) *ComissionValueRepository {
	return &ComissionValueRepository{
		db: db,
	}
}