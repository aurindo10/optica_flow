package orderrepository

import (
	"context"
	"optica_flow/internal/app/domain/orders"
	database "optica_flow/internal/app/infra/database/queries"
)

type OrderRepository struct {
	db *database.Queries
}

func (c *OrderRepository) Create(p *orders.Order) (*orders.Order, error) {
	params := database.CreateOrderParams{
		ID: p.ID,
		ClientID: p.ClientID,
		ProductName: p.ProductName,
		Quantity: p.Quantity,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		CompanyID: p.CompanyID,
		Fase: p.Fase,
	}
	result, error := c.db.CreateOrder(context.Background(), params)
	if error != nil {
		return nil, error
	}
	response := &orders.Order{
		ID: result.ID,
		ClientID: result.ClientID,
		ProductName: result.ProductName,
		Quantity: result.Quantity,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
		CompanyID: result.CompanyID,
		Fase: result.Fase,
		OrderDate: result.OrderDate,
	}
	return response, nil
}

func NewOrderRepository (db *database.Queries) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}