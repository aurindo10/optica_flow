package controllers

import (
	"errors"
	flowentries "optica_flow/internal/app/domain/flow_entries"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type FlowEntrieController struct {
	Create *flowentries.CreateFlowEntrie
	FindByDateRange *flowentries.FindByIntervalDate
	UpdateFlowEntrie *flowentries.UpdateFlowEntrie
	DeleteFlowEntrie *flowentries.DeleteFlowEntrie
}

func (r *FlowEntrieController) CreateFlowEntrie(c *fiber.Ctx) error {
	var request flowentries.CashFlowEntriesParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	if error :=  r.Validate(&request); error != nil {
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

func (r *FlowEntrieController) FindByRangeDate(c *fiber.Ctx) error {
	var request flowentries.FindByDateParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	if request.CompanyId == "" {
		return c.Status(fiber.StatusBadRequest).
		JSON(errors.New("o campo company id é obrigatório"))
	}
	if request.EndDate.IsZero() {
		return c.Status(fiber.StatusBadRequest).
		JSON(errors.New("o campo endDate id é obrigatório"))
	}
	if request.InitialDate.IsZero() {
		return c.Status(fiber.StatusBadRequest).
		JSON(errors.New("o campo endDate id é obrigatório"))
	}
	result, error := r.FindByDateRange.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).
		JSON(error.Error())
	}

	return c.Status(fiber.StatusOK).
	JSON(result)
}
func (r *FlowEntrieController) UpdateCashFlowEntrie(c *fiber.Ctx) error {
	var request flowentries.CashFlowEntriesUpdate
	if error := c.BodyParser(&request) ; error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	result, error := r.UpdateFlowEntrie.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).
	JSON(result)
}
func (r *FlowEntrieController) DeleteCashFlowEntrie(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(error.Error())
	}
	err := r.DeleteFlowEntrie.Execute(idParsed)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
		JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).
	JSON(fiber.Map{
		"message":"Deletado com sucesso",
	})
}
func NewFlowEntrieController(Create *flowentries.CreateFlowEntrie, FindByDateRange *flowentries.FindByIntervalDate,
	UpdateFlowEntrie *flowentries.UpdateFlowEntrie, DeleteFlowEntrie *flowentries.DeleteFlowEntrie,
	) *FlowEntrieController {
	return &FlowEntrieController{
		Create: Create,
		FindByDateRange: FindByDateRange,
		UpdateFlowEntrie: UpdateFlowEntrie,
		DeleteFlowEntrie: DeleteFlowEntrie,
	}
}

func (r *FlowEntrieController) Validate(p*flowentries.CashFlowEntriesParams) error {
	if p.Amount == 0.0 {
		return errors.New("o campo amount deve ser maior que 0")
	}
	if p.CompanyID == "" {
		return errors.New("o campo companyId deve é obrigatório")
	}
	if p.Description == ""{
		return errors.New("o campo descrição é obrigatório")
	}
	if p.WhoCreatedID == "" {
		return errors.New("o campo WhoCreatedID é obrigatório")

	}
	if p.WhoUpdatedID == "" {
		return errors.New("o campo WhoUpdatedID é obrigatório")

	}
	return nil
}