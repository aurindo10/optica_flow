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
func (c *ClientRepository) Find(id string) ([]*client.Client, error){
	result, error := c.db.FindClientsByCompanyid(context.Background(), id)
	if error != nil {
		return nil, error
	}
	response := c.mapResult(&result)
	return response, nil
}
func NewCLientRepository(db *database.Queries) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (c *ClientRepository) mapResult(result *[]database.Client) []*client.Client {
	var clients []*client.Client
	for _, v := range *result {
		clients = append(clients, &client.Client{
			ID: v.ID,
			FullName: v.FullName,
			Telefone: v.Telefone,
			Cpf: v.Cpf,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Email: v.Email,
			BirthDate: v.BirthDate,
			Adress: v.Adress,
			Gender: v.Gender,
			City: v.City,
			SellerID: v.SellerID,
			CompanyID: v.CompanyID,
			WhoCreatedID: v.WhoCreatedID,
			WhoUpdatedID: v.WhoUpdatedID,
		})
	}
	return clients
}