package product

import (
	"fmt"
)



type GetAllProducts struct {
	repository Repository
}

func (c *GetAllProducts) Execute(id string)([]*Product, error) {
	products, err :=  c.repository.GetAllProducts(id)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}
	return products, nil
}

func NewGetAllProducts(repository Repository) *GetAllProducts {
	return &GetAllProducts{repository}
}