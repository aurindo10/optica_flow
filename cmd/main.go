package main

import (
	"optica_flow/internal/app/domain/product"
	"optica_flow/internal/app/infra/database"
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
		fx.Invoke(func(*fasthttp.Server){}),
	)
	app.Run()
}