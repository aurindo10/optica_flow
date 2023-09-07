package fornecedorrepository

import (
	"context"
	"optica_flow/internal/app/domain/fornecedor"
	database "optica_flow/internal/app/infra/database/queries"
)


type FornecedorRepository struct {
	db *database.Queries
}

func (f *FornecedorRepository) CreateFornecedor(params *fornecedor.Fornecedor) (*fornecedor.Fornecedor, error) {
	ctx := context.Background()

	result, err := f.db.CreateFornecedor(ctx, database.CreateFornecedorParams{
		ID: params.ID,
		Name: params.Name,
		Email: params.Email,
		Telefone: params.Telefone,
		Adress: params.Adress,
		CompanyID: params.CompanyID,
		WhoCreatedID: params.WhoCreatedID,
		WhoUpdatedID: params.WhoUpdatedID,
		Cnpj: params.Cnpj,
	})
	if err != nil {
		return nil, err
	}
	response := &fornecedor.Fornecedor{
		ID: result.ID,
		Name: result.Name,
		Email: result.Email,
		Telefone: result.Telefone,
		Adress: result.Adress,
		CompanyID: result.CompanyID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
		Cnpj: result.Cnpj,
	}
	return response, nil
}
func NewFornecedorRepository(db *database.Queries) *FornecedorRepository {
	return &FornecedorRepository{
		db: db,
	}
}