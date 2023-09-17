package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)


type FlowBalanceRouter struct {
  controllers *controllers.CashFlowBalanceController
}

func (c *FlowBalanceRouter) Load(r *fiber.App) {
	r.Post("/flowbalance", c.controllers.CreateFlowBalance)
}

func NewFlowBalanceRouter(controllers *controllers.CashFlowBalanceController) *FlowBalanceRouter {
	return &FlowBalanceRouter{
		controllers: controllers,
	}
}