package fornecedor

import "github.com/google/uuid"


type Delete struct {
	repository Repository
}

func (c *Delete) Execute(id uuid.UUID) error {
	err := c.repository.Delete(id)
	if err != nil {
		return err	
	}
	return nil
}

func NewDeleleteFornecedor(repository Repository) *Delete {
	return &Delete{repository: repository}
}