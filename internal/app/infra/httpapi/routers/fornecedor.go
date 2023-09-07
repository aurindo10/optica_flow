package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type FornecedorRouter struct {
	controller *controllers.FornecedorController
}
func (c *FornecedorRouter) Load(r *fiber.App) {
	r.Post("/fornecedor", c.controller.CreateFornecedor)
}
func NewFornecedorRouter(controller *controllers.FornecedorController) *FornecedorRouter{
	return &FornecedorRouter{
		controller: controller,
	}
}