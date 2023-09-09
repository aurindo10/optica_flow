package client

import "github.com/google/uuid"

type FindOne struct {
	repository Repository
}
func (c *FindOne) Execute(id uuid.UUID) (*Client, error) {
	result, error := c.repository.FindOne(id)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewFindOne(repository Repository) *FindOne {
	return &FindOne{
		repository: repository,
	}
}