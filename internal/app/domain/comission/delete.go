package comission

import (
	"github.com/google/uuid"
)


type DeleteComission struct {
	repository Repository
}

func (c *DeleteComission) Execute(id uuid.UUID) error {
	error := c.repository.DeleteById(id)
	if error != nil {
		return error
	}
	return nil
}

func NewDeleteComission(repository Repository) *DeleteComission {
	return &DeleteComission{
		repository: repository,
	}
}