package orders

import "github.com/google/uuid"

type DeleteOrder struct {
	repository Repository
}

func (c *DeleteOrder) Execute(id uuid.UUID) error {
	err := c.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
func NewDeleteOrder(repository Repository) *DeleteOrder {
	return &DeleteOrder{repository}
}