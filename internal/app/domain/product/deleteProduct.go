package product

import "github.com/google/uuid"


type DeleteProduct struct {
	repository Repository
}


func (c *DeleteProduct) Execute(id uuid.UUID) error {
	error := c.repository.DeleteProduct(id)
	if error != nil {
		return error
	}	
	return nil
}

func NewDeleteProduct(repository Repository)  *DeleteProduct {
	return &DeleteProduct{repository}
}