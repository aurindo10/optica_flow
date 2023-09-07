package fornecedor


type CreateFornecedor struct {
	repository Repository
}

func (c *CreateFornecedor) Execute(request *Fornecedor) (*Fornecedor, error) {
	info := NewFornecedor(request)
	fornecedor, err := c.repository.CreateFornecedor(info)
	if err != nil {
		return nil, err
	}
	return fornecedor, nil
}

func NewCreateFornecedor (repository Repository) *CreateFornecedor {
	return &CreateFornecedor{repository: repository}
}