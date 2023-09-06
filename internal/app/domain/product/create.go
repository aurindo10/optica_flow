package product

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateProduct struct {
	repository Repository
}
type  CreateProductParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Fornecedor  string    `json:"fornecedor"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *CreateProduct) Execute(request *CreateProductParams) (*Product, error) {
	var info  = Params{
		ID:          request.ID,
		Name:        request.Name,
		Price:       request.Price,
		Fornecedor:  request.Fornecedor,
		Description: request.Description,
		Brand:       request.Brand,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}
	product := NewProduct(&info)
	err := c.repository.CreateProduct(product)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}
func NewCreateProduct(repository Repository) *CreateProduct {
	return &CreateProduct{repository}
}