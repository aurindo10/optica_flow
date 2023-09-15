package comissionvalue

import "github.com/google/uuid"

type DeleteComissionValue struct {
	repository Repository
}

func (c *DeleteComissionValue) Execute (id uuid.UUID) error {
	error := c.repository.DeleteComission(id)
	if error != nil {
		return error 
	}
	return nil
}

func NewDeleteComissionVAlue(repository Repository) *DeleteComissionValue {
	return &DeleteComissionValue{
		repository: repository,
	}
}