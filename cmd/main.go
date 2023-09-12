package main

import (
	"optica_flow/internal/app/domain/client"
	"optica_flow/internal/app/domain/fornecedor"
	"optica_flow/internal/app/domain/orders"
	"optica_flow/internal/app/domain/points"
	"optica_flow/internal/app/domain/product"
	productorder "optica_flow/internal/app/domain/product_order"
	"optica_flow/internal/app/infra/database"
	clientrepository "optica_flow/internal/app/infra/database/clientRepository"
	fornecedorrepository "optica_flow/internal/app/infra/database/fornecedorRepository"
	orderrepository "optica_flow/internal/app/infra/database/orderRepository"
	pointsrepository "optica_flow/internal/app/infra/database/pointsRepository"
	productorderrepository "optica_flow/internal/app/infra/database/productOrderRepository"
	"optica_flow/internal/app/infra/httpapi"
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"
	"optica_flow/internal/app/infra/httpapi/middleware"
	"optica_flow/internal/app/infra/httpapi/routers"

	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		httpapi.Module,
		database.Module,
		product.Module,
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
		fx.Invoke(func(*fasthttp.Server){}),
	)
	app.Run()
}