package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type ProductOrderRouter struct{
	controller *controllers.ProductOrderController
}

func (c *ProductOrderRouter) Load(r *fiber.App) {
	r.Post("/productorder", c.controller.Create)
}

func NewProductOrderRouter(controller *controllers.ProductOrderController) *ProductOrderRouter{
	return &ProductOrderRouter{
		controller: controller,
	}
}