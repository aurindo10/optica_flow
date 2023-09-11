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
	r.Get("/productorder/order/:id", c.controller.FindByOrderId)
	r.Put("/productorder", c.controller.UpdateProductOrderById)
}

func NewProductOrderRouter(controller *controllers.ProductOrderController) *ProductOrderRouter{
	return &ProductOrderRouter{
		controller: controller,
	}
}