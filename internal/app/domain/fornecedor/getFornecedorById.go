package fornecedor

import (
	"fmt"

	"github.com/google/uuid"
)

type GetFornecedorById struct {
	repository Repository
}
func (c *GetFornecedorById) Execute(id uuid.UUID) (*Fornecedor, error) {
	fmt.Printf("GetFornecedorById: %v\n", id)
	response, error := c.repository.GetFornecedorById(id)
	if error != nil {
		return nil, error
	}
	return response, nil
}
func NewGetFornecedorById(repository Repository) *GetFornecedorById {
	return &GetFornecedorById{repository: repository}
}