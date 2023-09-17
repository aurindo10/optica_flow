package main

import (
	cashflowout "optica_flow/internal/app/domain/cash_flow_out"
	"optica_flow/internal/app/domain/client"
	"optica_flow/internal/app/domain/comission"
	comissionvalue "optica_flow/internal/app/domain/comission_value"
	flowentries "optica_flow/internal/app/domain/flow_entries"
	"optica_flow/internal/app/domain/fornecedor"
	"optica_flow/internal/app/domain/orders"
	"optica_flow/internal/app/domain/points"
	"optica_flow/internal/app/domain/product"
	productorder "optica_flow/internal/app/domain/product_order"
	tradeproduct "optica_flow/internal/app/domain/trade_product"
	"optica_flow/internal/app/infra/database"
	clientrepository "optica_flow/internal/app/infra/database/clientRepository"
	comissionrepository "optica_flow/internal/app/infra/database/comissionRepository"
	comissionvaluerepository "optica_flow/internal/app/infra/database/comissionValueReposiitory"
	flowbalancerepository "optica_flow/internal/app/infra/database/flowBalanceRepository"
	flowentriesrepository "optica_flow/internal/app/infra/database/flowEntriesRepository"
	fornecedorrepository "optica_flow/internal/app/infra/database/fornecedorRepository"
	orderrepository "optica_flow/internal/app/infra/database/orderRepository"
	pointsrepository "optica_flow/internal/app/infra/database/pointsRepository"
	productorderrepository "optica_flow/internal/app/infra/database/productOrderRepository"
	tradeproductrepository "optica_flow/internal/app/infra/database/tradeProductRepository"
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
		tradeproduct.Module,
		tradeproductrepository.Module,
		comission.Module,
		comissionrepository.Module,
		comissionvalue.Module,
		comissionvaluerepository.Module,
		flowentries.Module,
		flowentriesrepository.Module,
		cashflowout.Module,
		flowbalancerepository.Module,
		fx.Invoke(func(*fasthttp.Server){}),
	)
	app.Run()
}