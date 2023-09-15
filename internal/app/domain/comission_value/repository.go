package comissionvalue

import "github.com/google/uuid"


type Repository interface {
	Create(*ComissionValues) (*ComissionValues, error)
	FindAll(id string) ([]*ComissionValues, error)
	DeleteComission(id uuid.UUID) error
}