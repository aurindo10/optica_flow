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
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	FornecedorID *string   `json:"fornecedor_id"`
	Description  string    `json:"description"`
	Brand        string    `json:"brand"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	BarCode      string    `json:"bar_code"`
	Quantity     int32     `json:"quantity"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}

func (c *CreateProduct) Execute(request *CreateProductParams) (*Product, error) {
	var info  = Params{
		ID:          request.ID,
		Name:        request.Name,
		Price:       request.Price,
		FornecedorID: request.FornecedorID,
		Description: request.Description,
		Brand:       request.Brand,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		BarCode:     request.BarCode,
		Quantity:    request.Quantity,
		CompanyID:   request.CompanyID,
		WhoCreatedID:request.WhoCreatedID,
		WhoUpdatedID:request.WhoUpdatedID,
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