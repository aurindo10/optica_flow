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
	r.Get("/point/seller/:id", c.controller.FindBySeller)
	r.Delete("/point/:id", c.controller.DeletePoints)
}
func NewPointsRouter(controller *controllers.PointsController) *PointsRouter{
	return &PointsRouter{
		controller: controller,
	}
}