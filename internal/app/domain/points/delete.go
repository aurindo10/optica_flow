package points

import "github.com/google/uuid"

type DeletePoints struct {
	repository Repository
}

func (c *DeletePoints) Execute(id uuid.UUID) error {
	error := c.repository.DeletePoints(id)
	if error != nil {
		return error
	}
	return nil
}
func NewDeletePoints(repository Repository) *DeletePoints {
	return &DeletePoints{
		repository: repository,
	}
}