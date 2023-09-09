package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type OrderRouter struct {
	constroller *controllers.OrderController
}

func (c *OrderRouter) Load(r *fiber.App) {
	r.Post("/order", c.constroller.CreateOrder)
	r.Get("/order/:id", c.constroller.FindOneOrder)
}

func NewOrderRouter(constroller *controllers.OrderController) *OrderRouter {
	return &OrderRouter{
		constroller: constroller,
	}
}