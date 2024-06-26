package fornecedor

import "github.com/google/uuid"

type Repository interface {
	CreateFornecedor(fornecedor *Fornecedor) (*Fornecedor, error)
	GetFornecedorById(id uuid.UUID) (*Fornecedor, error)
	FindAllFornecedoresByCompanyId(id string) ([]*Fornecedor, error)
	Update(fornecedor *FornecedorToUpdate) (*Fornecedor, error)
	Delete(id uuid.UUID) error
}
