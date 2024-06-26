package orders

import "github.com/google/uuid"


type Repository interface {
	Create (*Order) (*Order, error)
	FindOne (id uuid.UUID) (*Order, error)
	Update (*OrderToUpdate) (*Order, error)
	FindAll (companyId string) ([]*Order, error)
	Delete (id uuid.UUID) error
}