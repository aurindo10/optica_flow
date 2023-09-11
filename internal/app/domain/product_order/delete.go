package productorder

import "github.com/google/uuid"


type DeleteProductOrder struct {
	repository Repository
}


func (c *DeleteProductOrder) Execute(id uuid.UUID) error {
	error := c.repository.DeleteProductOrder(id)
	if error != nil {
		return error
	}
	return nil
}

func NewDeleteProductOrder(repository Repository) *DeleteProductOrder {
	return &DeleteProductOrder{
		repository: repository,
	}
}