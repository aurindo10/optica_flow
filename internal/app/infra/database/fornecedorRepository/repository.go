package fornecedorrepository

import (
	"context"
	"optica_flow/internal/app/domain/fornecedor"
	database "optica_flow/internal/app/infra/database/queries"

	"github.com/google/uuid"
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
func (c * FornecedorRepository) GetFornecedorById(id uuid.UUID) (*fornecedor.Fornecedor, error) {
	p, error := c.db.GetFornecedorByID(context.Background(), id)
	if error != nil {
		return nil, error
	}
	response := &fornecedor.Fornecedor{
		ID: p.ID,
		Name: p.Name,
		Email: p.Email,
		Telefone: p.Telefone,
		Adress: p.Adress,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		Cnpj: p.Cnpj,
	}
	return response, nil
}
func (c * FornecedorRepository) FindAllFornecedoresByCompanyId(id string) ([]*fornecedor.Fornecedor, error) {
	r, error := c.db.FindAllFornecedores(context.Background(), id)
	if error != nil {
		return nil, error
	}
	result := c.mapResult(&r)
	return result, nil 
}
func NewFornecedorRepository(db *database.Queries) *FornecedorRepository {
	return &FornecedorRepository{
		db: db,
	}
}
func (c *FornecedorRepository) mapResult(result *[]database.Fornecedor) []*fornecedor.Fornecedor {
	var fornecedores []*fornecedor.Fornecedor
	for _, v := range *result {
		fornecedores = append(fornecedores, &fornecedor.Fornecedor{
			ID: v.ID,
			Name: v.Name,
			Email: v.Email,
			Telefone: v.Telefone,
			Adress: v.Adress,
			CompanyID: v.CompanyID,
			WhoCreatedID: v.WhoCreatedID,
			WhoUpdatedID: v.WhoUpdatedID,
			Cnpj: v.Cnpj,
		})
	}
	return fornecedores
}

func (c *FornecedorRepository) Update(params *fornecedor.FornecedorToUpdate) (*fornecedor.Fornecedor, error){
	p := database.UpdateFornecedorParams{
		ID: params.ID,
		Name: params.Name,
		Email: params.Email,
		Telefone: params.Telefone,
		Adress: params.Adress,
		Cnpj: params.Cnpj,
	}
	result, error := c.db.UpdateFornecedor(context.Background(), p)
	if error != nil {
		return nil, error
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
func (c *FornecedorRepository) Delete(id uuid.UUID) error {
	error := c.db.DeleteFornecedorById(context.Background(), id)
	if error != nil {
		return error
	}
	return nil
}