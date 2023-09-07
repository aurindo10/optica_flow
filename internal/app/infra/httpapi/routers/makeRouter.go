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
) *fiber.App {
	r := fiber.New()
	FornecedorRouter.Load(r)
	ProductRouter.Load(r)
	return r
}