package product

import "github.com/google/uuid"

type GetProduct struct {
	repository Repository
}

func (p *GetProduct) Execute(id uuid.UUID) (*Product, error) {
	product, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}


func NewGetProduct(repository Repository) *GetProduct {
	return &GetProduct{repository: repository}
}