package product

import (
	"github.com/google/uuid"
)

type ProductRequestToUpdate struct {
	FornecedorID *string   `json:"fornecedor_id"`
	Name         *string    `json:"name"`
	Price        *float64   `json:"price"`
	Description  *string    `json:"description"`
	Brand        *string    `json:"brand"`
	BarCode      *string    `json:"bar_code"`
	Quantity     *int32     `json:"quantity"`
	WhoUpdatedID *string   `json:"who_updated_id"`
	ID           uuid.UUID `json:"id"`
}

type Repository interface {
	CreateProduct(product *Product) error
	GetAllProducts(id string) ([]*Product, error)
	UpdateProduct(product *ProductRequestToUpdate) (*Product, error)
	DeleteProduct(id uuid.UUID) error
	GetProductById(id uuid.UUID) (*Product, error)
}

