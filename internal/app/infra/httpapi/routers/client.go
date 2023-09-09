package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type ClientRouter struct {
	controller *controllers.ClientController
}

func (c *ClientRouter) Load(r *fiber.App) {
	r.Post("/client", c.controller.Create)
	r.Get("/client/company/:id", c.controller.Find)	
	r.Put("/client", c.controller.Update)
	r.Get("/client/:id", c.controller.FindOne)
}

func NewClientRouter(controller *controllers.ClientController) *ClientRouter{
	return &ClientRouter{
		controller: controller,
	}
}