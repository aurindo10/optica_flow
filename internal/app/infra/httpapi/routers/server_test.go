package routers_test

import (
	"bytes"
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

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)
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
		res, err := server.Test(req)
		if err != nil {
			t.Logf("Error: %v", err)
		}
		resBodyBytes, err := io.ReadAll(res.Body)
		t.Log(string(resBodyBytes))
		if err != nil {
			t.Logf("Error: %v", err)
		}
	})
}

