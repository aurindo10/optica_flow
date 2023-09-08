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
}

func NewClientRouter(controller *controllers.ClientController) *ClientRouter{
	return &ClientRouter{
		controller: controller,
	}
}