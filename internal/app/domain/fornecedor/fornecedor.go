package fornecedor

import (
	"github.com/google/uuid"
)

type Fornecedor struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Telefone     string    `json:"telefone"`
	Email        string    `json:"email"`
	Adress       string    `json:"adress"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	Cnpj         string    `json:"cnpj"`
}

func NewFornecedor(request *Fornecedor) *Fornecedor {
	return &Fornecedor{
		ID:           uuid.New(),
		Name:         request.Name,
		Telefone:     request.Telefone,
		Email:        request.Email,
		Adress:       request.Adress,
		CompanyID:    request.CompanyID,
		WhoCreatedID: request.WhoCreatedID,
		WhoUpdatedID: request.WhoUpdatedID,
		Cnpj:         request.Cnpj,
	}
}