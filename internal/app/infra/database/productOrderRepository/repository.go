package productorderrepository

import (
	"context"
	productorder "optica_flow/internal/app/domain/product_order"
	database "optica_flow/internal/app/infra/database/queries"

	"github.com/google/uuid"
)


type ProductOrderRepository struct {
	db *database.Queries
}

func (c *ProductOrderRepository) Create(p *productorder.ProductOrder) (*productorder.ProductOrder, error) {
	params := database.CreateProductOrderParams{
		ID: p.ID,
		ProductID: p.ProductID,
		OrderID: p.OrderID,
		Amout: p.Amout,
	}
	result, error := c.db.CreateProductOrder(context.Background(), params)
	if error != nil {
		return nil, error
	}
	response := &productorder.ProductOrder{
		ID: result.ID,
		ProductID: result.ProductID,
		OrderID: result.OrderID,
		Amout: result.Amout,
	}
	return response, nil
}
func (c *ProductOrderRepository) FindByOrderId(id *string) ([]*productorder.ProductOrder, error){
	result, err := c.db.FindAllProductOrdersByOrderId(context.Background(), id)
	if err != nil {
		return nil, err
	}
	var response []*productorder.ProductOrder
	for _, v := range result {
		response = append(response, &productorder.ProductOrder{
			ID: v.ID,
			ProductID: v.ProductID,
			OrderID: v.OrderID,
			Amout: v.Amout,
		})
	}
	return response, nil
}
func (c *ProductOrderRepository) UpdateProductOrder(p *productorder.ProductOrderToUpdate) (*productorder.ProductOrder, error) {
	params := database.UpdateProductOrderParams{
		ID: p.ID,
		Amout: p.Amout,
	}
	result, error := c.db.UpdateProductOrder(context.Background(), params)
	if error != nil {
		return nil, error
	}
	response := &productorder.ProductOrder{
		ID: result.ID,
		ProductID: result.ProductID,
		OrderID: result.OrderID,
		Amout: result.Amout,
	}
	return response, nil
}
func (c *ProductOrderRepository) DeleteProductOrder(id uuid.UUID) error {
	error := c.db.DeleteProductOrderById(context.Background(), id)
	if error != nil {
		return error
	}
	return nil
}
func NewProductOrderRepository(db *database.Queries) *ProductOrderRepository {
	return &ProductOrderRepository{
		db: db,
	}
}