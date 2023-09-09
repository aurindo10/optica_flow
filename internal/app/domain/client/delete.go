package client

import (
	"github.com/google/uuid"
)


type DeleteClient struct {
	repository Repository
}

func (c *DeleteClient) Execute(id uuid.UUID) error {
	error := c.repository.DeleteOne(id)
	if error != nil {
		return error
	}
	return nil 
}

func NewDeleteClient(repository Repository) *DeleteClient {
	return &DeleteClient{repository: repository}
}