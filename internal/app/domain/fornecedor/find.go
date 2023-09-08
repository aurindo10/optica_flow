package fornecedor

type Find struct {
	repository Repository
}
func (r *Find) Execute(id string) ([]*Fornecedor, error) {
	response, error := r.repository.FindAllFornecedoresByCompanyId(id)
	if error != nil {
		return nil, error
	}
	return response, nil
}
func NewFind(repository Repository) *Find {
	return &Find{repository: repository}
}