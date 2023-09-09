package orders

import "github.com/google/uuid"

type FindOne struct {
	repository Repository
}

func (c *FindOne) Execute(id uuid.UUID) (*Order, error) {
	result, err := c.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewFindOne(repository Repository) *FindOne {
	return &FindOne{
		repository: repository,
	}
}

