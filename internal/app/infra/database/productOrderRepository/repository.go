package productorderrepository

import (
	"context"
	productorder "optica_flow/internal/app/domain/product_order"
	database "optica_flow/internal/app/infra/database/queries"
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

func NewProductOrderRepository(db *database.Queries) *ProductOrderRepository {
	return &ProductOrderRepository{
		db: db,
	}
}