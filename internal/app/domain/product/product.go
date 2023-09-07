package product

import (
	"time"

	"github.com/google/uuid"
)


type Params struct {
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
type Product struct {
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

func NewProduct(request * Params) *Product {
	return &Product{
		ID          :uuid.New(),
		Name        :request.Name,  
		Price       :request.Price, 
		FornecedorID  : request.FornecedorID,  
		Description : request.Description,  
		Brand       : request.Brand,  
		CreatedAt   :time.Now(),
		UpdatedAt   :time.Now(),
		BarCode     :request.BarCode,
		Quantity    :request.Quantity,
		CompanyID   :request.CompanyID,
		WhoCreatedID:request.WhoCreatedID,
		WhoUpdatedID:request.WhoUpdatedID,
	}
}

