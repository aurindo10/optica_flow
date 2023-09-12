package pointsrepository

import (
	"context"
	"fmt"
	"optica_flow/internal/app/domain/points"
	database "optica_flow/internal/app/infra/database/queries"
)


type PointsRepository struct {
	db *database.Queries
}

func (c *PointsRepository) Create(p *points.Points) (*points.Points, error){
	fmt.Printf("GetFornecedorById: %v\n", p.OrderID)
	request := database.CreatePointsParams{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Active: p.Active,
		Ammount: p.Ammount,
		ValidDate: p.ValidDate,
		CompanyID: p.CompanyID,
		OrderID: p.OrderID,
	}
	result, error := c.db.CreatePoints(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &points.Points{
		ID: result.ID,
		Name: result.Name,
		Description: result.Description,
		Active: result.Active,
		Ammount: result.Ammount,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		ValidDate: result.ValidDate,
		CompanyID: result.CompanyID,
		OrderID: result.OrderID,
	}
	return response, nil
}

func NewPointsRepository(db *database.Queries) *PointsRepository {
	return &PointsRepository{
		db: db,
	}
}