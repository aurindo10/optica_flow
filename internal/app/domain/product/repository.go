package product

import (
	"time"

	"github.com/google/uuid"
)

type ProductRequestToUpdate struct {
	ID 		uuid.UUID `json:"id"`
	Name 	*string `json:"name"`
	Price 	*float64 `json:"price"`
	Fornecedor *string `json:"fornecedor"`
	Description *string `json:"description"`
	Brand 		*string `json:"brand"`
	UpdateAt 	time.Time `json:"update_at"`
}

type Repository interface {
	CreateProduct(product *Product) error
	GetAllProducts() ([]*Product, error)
	UpdateProduct(product *ProductRequestToUpdate) (*Product, error)
	DeleteProduct(id uuid.UUID) error
}

