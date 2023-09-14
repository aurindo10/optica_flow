package routers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"optica_flow/internal/app/domain/client"
	"optica_flow/internal/app/domain/comission"
	"optica_flow/internal/app/domain/fornecedor"
	"optica_flow/internal/app/domain/orders"
	"optica_flow/internal/app/domain/points"
	"optica_flow/internal/app/domain/product"
	productorder "optica_flow/internal/app/domain/product_order"
	tradeproduct "optica_flow/internal/app/domain/trade_product"
	"optica_flow/internal/app/infra/database"
	clientrepository "optica_flow/internal/app/infra/database/clientRepository"
	comissionrepository "optica_flow/internal/app/infra/database/comissionRepository"
	fornecedorrepository "optica_flow/internal/app/infra/database/fornecedorRepository"
	orderrepository "optica_flow/internal/app/infra/database/orderRepository"
	pointsrepository "optica_flow/internal/app/infra/database/pointsRepository"
	productorderrepository "optica_flow/internal/app/infra/database/productOrderRepository"
	tradeproductrepository "optica_flow/internal/app/infra/database/tradeProductRepository"
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"
	"optica_flow/internal/app/infra/httpapi/middleware"
	"optica_flow/internal/app/infra/httpapi/routers"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)



type Client struct {
	ID           uuid.UUID `json:"id"`
	FullName     string    `json:"full_name"`
	Telefone     string    `json:"telefone"`
	Cpf          string   `json:"cpf"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birth_date"`
	Adress       string   `json:"adress"`
	Gender       string   `json:"gender"`
	City         string   `json:"city"`
	SellerID     string   `json:"seller_id"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}
func TestClient(t *testing.T) {
	var server *fiber.App
	app := fxtest.New(t,
		product.Module,
		database.TestModule,
		routers.Module,
		controllers.Module,
		middleware.Module,
		fornecedor.Module,
		client.Module,
		fornecedorrepository.Module,
		clientrepository.Module,
		orders.Module,
		orderrepository.Module,
		productorder.Module,
		productorderrepository.Module,
		points.Module,
		pointsrepository.Module,
		tradeproduct.Module,
		tradeproductrepository.Module,
		comission.Module,
		comissionrepository.Module,
		fx.Populate(&server),
	)

	app.RequireStart()
	defer app.RequireStop()
	t.Run("Test Create Client", func(t *testing.T) {
		reqBody := `{
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
		  }`
		req := httptest.NewRequest("POST", "/client", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		res, err := server.Test(req)
		if err != nil {
			t.Logf("Error: %v", err)
		}
		resBodyBytes, err := io.ReadAll(res.Body)
		var resClient, reqClient Client
		json.Unmarshal([]byte(reqBody), &reqClient)
		json.Unmarshal(resBodyBytes, &resClient)
		if reqClient.FullName != "" && reqClient.FullName != resClient.FullName {
			t.Errorf("FullName does not match")
		}
		if reqClient.Telefone != "" && reqClient.Telefone != resClient.Telefone {
			t.Errorf("Telefone does not match")
		}
		if reqClient.Email != "" && reqClient.Email != resClient.Email {
			t.Errorf("Email does not match")
		}
		if reqClient.Cpf != "" && reqClient.Cpf != resClient.Cpf {
			t.Errorf("CPF does not match")
		}
		if reqClient.BirthDate != resClient.BirthDate {
			t.Errorf("BirthDate does not match")
		}
		if reqClient.Adress != "" && reqClient.Adress != resClient.Adress {
			t.Errorf("Adress does not match")
		}
		if reqClient.Gender != "" && reqClient.Gender != resClient.Gender {
			t.Errorf("Gender does not match")
		}
		if reqClient.City != "" && reqClient.City != resClient.City {
			t.Errorf("City does not match")
		}
		if reqClient.SellerID != "" && reqClient.SellerID != resClient.SellerID {
			t.Errorf("SellerID does not match")
		}
		if reqClient.CompanyID != "" && reqClient.CompanyID != resClient.CompanyID {
			t.Errorf("CompanyID does not match")
		}
		if reqClient.WhoCreatedID != "" && reqClient.WhoCreatedID != resClient.WhoCreatedID {
			t.Errorf("WhoCreatedID does not match")
		}
		if reqClient.WhoUpdatedID != "" && reqClient.WhoUpdatedID != resClient.WhoUpdatedID {
			t.Errorf("WhoUpdatedID does not match")
		}
		if err != nil {
			t.Logf("Error: %v", err)
		}
	})	
}

