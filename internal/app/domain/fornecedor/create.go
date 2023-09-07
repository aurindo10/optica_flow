package fornecedor


type CreateFornecedor struct {
	repository Repository
}

func (c *CreateFornecedor) Execute(request *Fornecedor) (*Fornecedor, error) {
	p := NewFornecedor(request)
	result := &Fornecedor{
		ID: 		 p.ID,
		Name: 		 p.Name,
		Telefone: 	 p.Telefone,
		Email: 		 p.Email,
		Adress: 	 p.Adress,
		CompanyID: 	 p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		Cnpj: 		 p.Cnpj,
	}
	fornecedor, err := c.repository.CreateFornecedor(result)
	if err != nil { 
		return nil, err
	}
	return fornecedor, nil
}

func NewCreateFornecedor (repository Repository) *CreateFornecedor {
	return &CreateFornecedor{repository: repository}
}