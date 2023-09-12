package points

import "github.com/google/uuid"

type Repository interface {
	Create(*Points) (*Points, error)
	FindBySellerId(id string) ([]*Points, error)
	DeletePoints(id uuid.UUID) error
}