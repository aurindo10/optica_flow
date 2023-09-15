package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)


type FlowEntrieRouter struct {
	controller *controllers.FlowEntrieController
}

func (c *FlowEntrieRouter) Load(r *fiber.App) {
	r.Post("/flowentrie", c.controller.CreateFlowEntrie)
}

func NewFlowEntrieRouter(controller *controllers.FlowEntrieController) *FlowEntrieRouter {
	return &FlowEntrieRouter{
		controller: controller,
	}
}