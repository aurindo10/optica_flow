package client

import (
	"time"

	"github.com/google/uuid"
)



type UpdateClient struct {
	repository Repository
}

type ClientToUpdate struct {
	ID           uuid.UUID `json:"id"`
	FullName     *string    `json:"full_name"`
	Telefone     *string    `json:"telefone"`
	Cpf          *string   `json:"cpf"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Email        *string    `json:"email"`
	BirthDate    *time.Time `json:"birth_date"`
	Adress       *string   `json:"adress"`
	Gender       *string   `json:"gender"`
	City         *string   `json:"city"`
	SellerID     *string   `json:"seller_id"`
	CompanyID    *string    `json:"company_id"`
	WhoUpdatedID *string    `json:"who_updated_id"`
}
func (c *UpdateClient) Execute(client *ClientToUpdate) (*Client, error) {
	response, error := c.repository.Update(client)
	if error != nil {
		return nil, error
	}
	return response, nil
}
func NewUpdateClient(repository Repository) *UpdateClient {
	return &UpdateClient{repository: repository}
}