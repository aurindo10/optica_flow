package cashflowout

import (
	"github.com/google/uuid"
)

type DeleteCashFlowBalance struct {
	repository Repository
}

func (c *DeleteCashFlowBalance) Execute(id uuid.UUID) error {
	error := c.repository.Delete(id)
	if error != nil {
		return error
	}
	return nil
}

func NewDeleteCashFlowBalance(repository Repository) *DeleteCashFlowBalance {
	return &DeleteCashFlowBalance{
		repository: repository,
	}
}