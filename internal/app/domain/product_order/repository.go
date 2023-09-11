package productorder

import "github.com/google/uuid"

type Repository interface {
	Create(p *ProductOrder) (*ProductOrder, error)
	FindByOrderId(orderId *string) ([]*ProductOrder, error)
	UpdateProductOrder(p *ProductOrderToUpdate) (*ProductOrder, error)
	DeleteProductOrder(id uuid.UUID) error
}
