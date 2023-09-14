package routers

import (
	"bytes"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)


type TestClient struct {
	Error *error
	Res *int
}
func NewTestClient(app *fiber.App) *TestClient {
	println("olaa")
	req := httptest.NewRequest("POST", "/client", bytes.NewBufferString(`{
		"full_name": "Seu Nome",
		"telefone": "Seu Telefone",
		"email": "seu.email@exemplo.com",
		"cpf": "608.078.413-07",
		"birth_date": "2000-01-01T00:00:00Z",
		"adress": "Seu Endereço",
		"gender": "Seu Gênero",
		"city": "Sua Cidade",
		"seller_id": "ID do Vendedor",
		"company_id": "dasdasdasdas",
		"who_created_id": "ID de Quem Criou",
		"who_updated_id": "ID de Quem Atualizou"
	  }`))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		return &TestClient{
			Res: nil,
			Error: &err,
		}
	}
	return &TestClient{
		Res: &res.StatusCode,
		Error: nil,
	}
}
var TestClientModel = fx.Provide(
	NewTestClient,
)