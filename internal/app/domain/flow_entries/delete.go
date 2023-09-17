package flowentries

import "github.com/google/uuid"

type DeleteFlowEntrie struct {
	repository Repository
}

func (c *DeleteFlowEntrie) Execute(id uuid.UUID) error {
	error := c.repository.Delete(id)
	if error != nil {
		return error 
	}
	return nil
}

func NewDeleteFlowEntrie(repository Repository) *DeleteFlowEntrie {
	return &DeleteFlowEntrie{
		repository: repository,
	}
}