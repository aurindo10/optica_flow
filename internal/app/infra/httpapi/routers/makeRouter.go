package routers

import (
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Load()
}
func MakeRouter(
	ProductRouter *ProductRouter,
	FornecedorRouter *FornecedorRouter,
	ClientRouter *ClientRouter,
	OrderRouter *OrderRouter,
	ProduProductOrderRouter *ProductOrderRouter,
	PointsRouter *PointsRouter,
	TradeProductController *TradeProductController,
	ComissionController *ComissionRouter,
	ComissionValueRouter *ComissionValueRouter,
) *fiber.App {
	r := fiber.New()
	FornecedorRouter.Load(r)
	ProductRouter.Load(r)
	ClientRouter.Load(r)
	OrderRouter.Load(r)
	ProduProductOrderRouter.Load(r)
	PointsRouter.Load(r)
	TradeProductController.Load(r)
	ComissionController.Load(r)
	ComissionValueRouter.Load(r)
	return r
}