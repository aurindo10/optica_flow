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
	r.Get("/comission/user/:id", c.controller.FindByUser)
	r.Delete("/comission/:id", c.controller.DeleteComissionById)
}

func NewComissionRouter(controller *controllers.ComissionController) *ComissionRouter {
	return &ComissionRouter{
		controller: controller,
	}
}