package comissionrepository

import (
	"context"
	"optica_flow/internal/app/domain/comission"
	database "optica_flow/internal/app/infra/database/queries"

	"github.com/google/uuid"
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
func (c *ComissionRepository) FindByUserId(id string) ([]*comission.Commission, error){
	result, error := c.db.FindComissionByUserId(context.Background(), id)
	if error != nil {
		return nil, error
	}
	var comissions []*comission.Commission
	for _, oneComission := range result {
		comissions = append(comissions, &comission.Commission{
			ID: oneComission.ID,
			Name: oneComission.Name,
			Description: oneComission.Description,
			CreatedAt: oneComission.CreatedAt,
			UpdatedAt: oneComission.UpdatedAt,
			Value: oneComission.Value,
			CompanyID: oneComission.CompanyID,
			WhoCreatedID: oneComission.WhoCreatedID,
			WhoUpdatedID: oneComission.WhoUpdatedID,
			OrderID: oneComission.OrderID,
		})
	}
	return comissions, nil
}
func (c *ComissionRepository) DeleteById(id uuid.UUID) error {
	error := c.db.DeleteComissionById(context.Background(), id)
	if error != nil {
		return error
	}
	return nil
}
func NewComissionRepository( db *database.Queries) *ComissionRepository {
	return &ComissionRepository{
		db: db,
	}
}
