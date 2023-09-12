package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"

	"github.com/gofiber/fiber/v2"
)

type TradeProductController struct {
	controller *controllers.TradeProductController
}

func (c *TradeProductController) Load(r *fiber.App) {
	r.Post("/tradeproduct", c.controller.CreateTradeProduct)
	r.Get("/tradeproduct/company/:id", c.controller.FindAllTradeProducts)
}

func NewTradeProductController(controller *controllers.TradeProductController) *TradeProductController {
	return &TradeProductController{
		controller: controller,
	}
}