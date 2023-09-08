package client

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"full_name"`
	Telefone     string    `json:"telefone"`
	Cpf          *string   `json:"cpf"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birth_date"`
	Adress       *string   `json:"adress"`
	Gender       *string   `json:"gender"`
	City         *string   `json:"city"`
	SellerID     *string   `json:"seller_id"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}



func NewClient(resquest *Client) *Client {
	return &Client{
		ID:           uuid.New(),
		FullName:     resquest.FullName,
		Telefone:     resquest.Telefone,
		Cpf:          resquest.Cpf,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        resquest.Email,
		BirthDate:    resquest.BirthDate,
		Adress:       resquest.Adress,
		Gender:   		resquest.Cpf,
		City:         resquest.City,
		SellerID:     resquest.SellerID,
		CompanyID:    resquest.CompanyID,
		WhoCreatedID: resquest.WhoCreatedID,
		WhoUpdatedID: resquest.WhoUpdatedID,
	}
}