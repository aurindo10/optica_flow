package productorder

import "github.com/google/uuid"

type FindById struct {
	repository Repository
}

func (c *FindById) Execute(id uuid.UUID) (*ProductOrder, error) {
	return c.repository.FindById(id)
}

func NewFindByIdRequest(repository Repository) *FindById {
	return &FindById{repository: repository}
}