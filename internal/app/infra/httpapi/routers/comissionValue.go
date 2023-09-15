package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)


type ComissionValueRouter struct {
	controller *controllers.ComissionValueController
}

func (c *ComissionValueRouter) Load(r *fiber.App) {
	r.Post("/comissionvalue", c.controller.CreateValueComission)
}

func NewComissionValueRouter(controller *controllers.ComissionValueController) *ComissionValueRouter{
	return &ComissionValueRouter{
		controller: controller,
	}
}