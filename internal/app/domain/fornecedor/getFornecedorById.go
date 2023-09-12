package fornecedor

import (
	"github.com/google/uuid"
)

type GetFornecedorById struct {
	repository Repository
}
func (c *GetFornecedorById) Execute(id uuid.UUID) (*Fornecedor, error) {
	response, error := c.repository.GetFornecedorById(id)
	if error != nil {
		return nil, error
	}
	return response, nil
}
func NewGetFornecedorById(repository Repository) *GetFornecedorById {
	return &GetFornecedorById{repository: repository}
}