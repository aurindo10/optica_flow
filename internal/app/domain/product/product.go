package product

import (
	"time"

	"github.com/google/uuid"
)


type Params struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Fornecedor  string    `json:"fornecedor"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type Product struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Fornecedor  string    `json:"fornecedor"`
		Description string    `json:"description"`
		Brand       string    `json:"brand"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(request * Params) *Product {
	return &Product{
		ID          :uuid.New(),
		Name        :request.Name,  
		Price       :request.Price, 
		Fornecedor  : request.Fornecedor,  
		Description : request.Description,  
		Brand       : request.Brand,  
		CreatedAt   :time.Now(),
		UpdatedAt   :time.Now(),
	}
}

