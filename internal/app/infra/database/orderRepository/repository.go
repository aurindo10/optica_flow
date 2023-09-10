package orderrepository

import (
	"context"
	"optica_flow/internal/app/domain/orders"
	database "optica_flow/internal/app/infra/database/queries"

	"github.com/google/uuid"
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
func (c *OrderRepository) FindOne(id uuid.UUID) (*orders.Order, error){
	result, error := c.db.FindOneOrderById(context.Background(), id)
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
	if error != nil {
		return nil, error
	}
	return response, nil
}
func (c *OrderRepository ) Update(p *orders.OrderToUpdate) (*orders.Order, error){
	params := database.UpdateOneOrderParams{
		ID: p.ID,
		ProductName: p.ProductName,
		Quantity: p.Quantity,
		WhoUpdatedID: p.WhoUpdatedID,
		Fase: p.Fase,
	}
	result, error := c.db.UpdateOneOrder(context.Background(), params)
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
func (c *OrderRepository) FindAll(companyId string) ([]*orders.Order, error) {
	result, error := c.db.FindAllORdersByCompanyid(context.Background(), companyId)
	if error != nil {
		return nil, error
	}
	response := c.mapResult(&result)
	return response, nil
}
func (c *OrderRepository) Delete(id uuid.UUID) error {
	error := c.db.DeleteOrderById(context.Background(), id)
	if error != nil {
		return error
	}
	return nil
}
func NewOrderRepository (db *database.Queries) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}


func (c *OrderRepository) mapResult(result *[]database.Orders) []*orders.Order {
	var ordersResponse []*orders.Order
	for _, v := range *result {
		ordersResponse = append(ordersResponse, &orders.Order{
			ID: v.ID,
			ClientID: v.ClientID,
			ProductName: v.ProductName,
			Quantity: v.Quantity,
			WhoCreatedID: v.WhoCreatedID,
			WhoUpdatedID: v.WhoUpdatedID,
			CompanyID: v.CompanyID,
			Fase: v.Fase,
			OrderDate: v.OrderDate,
		})
	}
	return ordersResponse
}
