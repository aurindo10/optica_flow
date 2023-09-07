package fornecedorrepository

import (
	"context"
	"optica_flow/internal/app/domain/fornecedor"
	database "optica_flow/internal/app/infra/database/queries"
)


type FornecedorRepository struct {
	db *database.Queries
}

func (f *FornecedorRepository) CreateFornecedor(fornecedor *fornecedor.Fornecedor) (*database.Fornecedor, error) {
	ctx := context.Background()
	result, err := f.db.CreateFornecedor(ctx, database.CreateFornecedorParams{
		ID: fornecedor.ID,
		Name: fornecedor.Name,
		Email: fornecedor.Email,
		Telefone: fornecedor.Telefone,
		Adress: fornecedor.Adress,
		CompanyID: fornecedor.CompanyID,
		WhoCreatedID: fornecedor.WhoCreatedID,
		WhoUpdatedID: fornecedor.WhoUpdatedID,
		Cnpj: fornecedor.Cnpj,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewFornecedorRepository(db *database.Queries) *FornecedorRepository {
	return &FornecedorRepository{
		db: db,
	}
}