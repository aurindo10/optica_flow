package pointsrepository

import (
	"context"
	"optica_flow/internal/app/domain/points"
	database "optica_flow/internal/app/infra/database/queries"
)


type PointsRepository struct {
	db *database.Queries
}

func (c *PointsRepository) Create(p *points.Points) (*points.Points, error){
	request := database.CreatePointsParams{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Active: p.Active,
		Ammount: p.Ammount,
		ValidDate: p.ValidDate,
		CompanyID: p.CompanyID,
		OrderID: p.OrderID,
		SellerID: p.SellerID,
		
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
		SellerID: result.SellerID,
	}
	return response, nil
}
func (c *PointsRepository) FindBySellerId(id string) ([]*points.Points, error) {
	result, error := c.db.FindPointsBySellerId(context.Background(), id)
	if error != nil {
		return nil, error
	}
	var allPoints []*points.Points
	for _, one := range result {
		allPoints = append(allPoints, &points.Points{
			ID: one.ID,
			Name: one.Name,
			Description: one.Description,
			Active: one.Active,
			Ammount: one.Ammount,
			CreatedAt: one.CreatedAt,
			UpdatedAt: one.UpdatedAt,
			ValidDate: one.ValidDate,
			CompanyID: one.CompanyID,
			OrderID: one.OrderID,
			SellerID: one.SellerID,
		})
	}
	return allPoints, nil
}
func NewPointsRepository(db *database.Queries) *PointsRepository {
	return &PointsRepository{
		db: db,
	}
}