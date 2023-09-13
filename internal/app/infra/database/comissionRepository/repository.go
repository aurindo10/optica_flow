package comissionrepository

import (
	"context"
	"optica_flow/internal/app/domain/comission"
	database "optica_flow/internal/app/infra/database/queries"
)


type ComissionRepository struct {
	db *database.Queries
}

func (c *ComissionRepository) Create(p *comission.Commission) (*comission.Commission, error) {
	request := database.CreateComissionParams{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoCreatedID,
		OrderID: p.OrderID,
		Value: p.Value,
	}
	result, error := c.db.CreateComission(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &comission.Commission{
		ID: result.ID,
		Name: result.Name,
		Description: result.Description,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		CompanyID: result.CompanyID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoCreatedID,
		OrderID: result.OrderID,
		Value: result.Value,
	}
	return response, nil
}
func NewComissionRepository( db *database.Queries) *ComissionRepository {
	return &ComissionRepository{
		db: db,
	}
}
