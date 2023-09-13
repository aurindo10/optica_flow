package tradeproduct

import "github.com/google/uuid"


type DeleteTradeProduct struct {
	repository Repository
}

func (c *DeleteTradeProduct) Execute(id uuid.UUID) error {
	error := c.repository.Delete(id)
	if error != nil {
		return error
	}
	return nil
}

func NewDeleteTradeProduct(repository Repository) *DeleteTradeProduct {
	return &DeleteTradeProduct{
		repository: repository,
	}
}