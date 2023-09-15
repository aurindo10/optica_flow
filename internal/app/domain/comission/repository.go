package comission

import "github.com/google/uuid"


type Repository interface {
	Create(*Commission) (*Commission, error)
	FindByUserId(id string) ([]*Commission, error)
	DeleteById(id uuid.UUID) error
	UpdateById(*CommissionToUpdate) (*Commission, error)
}
