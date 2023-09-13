package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)


type ComissionRouter struct {
	controller *controllers.ComissionController
}

func (c *ComissionRouter) Load(r *fiber.App) {
	r.Post("/comission", c.controller.CreateComission)
}

func NewComissionRouter(controller *controllers.ComissionController) *ComissionRouter {
	return &ComissionRouter{
		controller: controller,
	}
}