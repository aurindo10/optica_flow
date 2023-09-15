package controllers

import (
	flowentries "optica_flow/internal/app/domain/flow_entries"

	"github.com/gofiber/fiber/v2"
)


type FlowEntrieController struct {
	Create *flowentries.CreateFlowEntrie
}

func (r *FlowEntrieController) CreateFlowEntrie(c *fiber.Ctx) error {
	var request flowentries.CashFlowEntriesParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	result, error := r.Create.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).
		JSON(error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func NewFlowEntrieController(Create *flowentries.CreateFlowEntrie) *FlowEntrieController {
	return &FlowEntrieController{
		Create: Create,
	}
}