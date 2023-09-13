package tradeproduct

import "github.com/google/uuid"

type Repository interface {
	Create(*TradeProduct) (*TradeProduct, error)
	FindAll(id string) ([]*TradeProduct, error)
	Update(*TradeProducToUpdate) (*TradeProduct, error)
	Delete(id uuid.UUID) error
}