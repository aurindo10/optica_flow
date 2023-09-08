package clientrepository

import (
	"context"
	"optica_flow/internal/app/domain/client"
	database "optica_flow/internal/app/infra/database/queries"
)

type ClientRepository struct {
	db *database.Queries
}

func (p * ClientRepository) Create(request *client.Client) (*client.Client, error){
		info := database.CreateClientParams{
			ID: request.ID,
			FullName: request.FullName,
			Telefone: request.Telefone,
			Cpf: request.Cpf,
			Email: request.Email,
			BirthDate: request.BirthDate,
			Adress: request.Adress,
			Gender: request.Gender,
			City: request.City,
			SellerID: request.SellerID,
			CompanyID: request.CompanyID,
			WhoCreatedID: request.WhoCreatedID,
			WhoUpdatedID: request.WhoUpdatedID,
		}
		result, error := p.db.CreateClient(context.Background(), info)
		if error != nil {
			return nil, error
		}
		response := client.Client{
				ID: result.ID,
				FullName: result.FullName,
				Telefone: result.Telefone,
				Cpf: result.Cpf,
				Email: result.Email,
				BirthDate: result.BirthDate,
				CreatedAt: result.CreatedAt,
				UpdatedAt: result.UpdatedAt,
				Adress: result.Adress,
				Gender: result.Gender,
				City: result.City,
				SellerID: result.SellerID,
				CompanyID: result.CompanyID,
				WhoCreatedID: result.WhoCreatedID,
				WhoUpdatedID: result.WhoUpdatedID,
			}
		return &response, nil 
}

func NewCLientRepository(db *database.Queries) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}