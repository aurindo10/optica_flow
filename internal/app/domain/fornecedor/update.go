package fornecedor


type Update struct {
	repository Repository
}

func (p *Update) Execute(fornecedor *FornecedorToUpdate) (*Fornecedor, error) {
	result, error := p.repository.Update(fornecedor)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewFornecedorUpdate(repository Repository) *Update {
	return &Update{repository: repository}
}