package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type PointsRouter struct {
	controller *controllers.PointsController
}
func (c *PointsRouter) Load(r *fiber.App) {
	r.Post("/point", c.controller.CreatePoints)
}
func NewPointsRouter(controller *controllers.PointsController) *PointsRouter{
	return &PointsRouter{
		controller: controller,
	}
}