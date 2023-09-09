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
) *fiber.App {
	r := fiber.New()
	FornecedorRouter.Load(r)
	ProductRouter.Load(r)
	ClientRouter.Load(r)
	OrderRouter.Load(r)
	return r
}