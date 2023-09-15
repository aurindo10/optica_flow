package controllers

import (
	"errors"
	comissionvalue "optica_flow/internal/app/domain/comission_value"

	"github.com/gofiber/fiber/v2"
)


type ComissionValueController struct {
	CreateComissionValue *comissionvalue.CreateValueComission
	FindComissionValue *comissionvalue.FindAllByCompanyId
}



func (r *ComissionValueController) CreateValueComission(c *fiber.Ctx) error {
	var comissionValue comissionvalue.ComissionValuesParams
	if error := c.BodyParser(&comissionValue); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{"error": error.Error()})
	}
	if error := r.Validate(&comissionValue); error != nil {
		return c.Status(fiber.StatusBadRequest).
		JSON(fiber.Map{"error": error.Error()})
	}
	result, error := r.CreateComissionValue.Execute(&comissionValue)
	if error != nil {
		c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (r *ComissionValueController) FindAllComissionValue(c *fiber.Ctx) error {
	id := c.Params("id")
	result, error := r.FindComissionValue.Execute(id)
	if error != nil {
		c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (r *ComissionValueController) Validate(request *comissionvalue.ComissionValuesParams) error {
	if request.WhoCreatedID == "" {
		return errors.New("whoCreatedID é obrigatório")
	}
	if request.WhoUpdatedID == "" {
		return errors.New("whoUpdatedID é obrigatório")
	}
	if request.Percentage == 0.0 {
		return errors.New("percentage não pode ser zero")
	}
	if request.Type == "" {
		return errors.New("type é obrigatório")
	}
	return nil
}

func NewComissionValueController(createComissionValue *comissionvalue.CreateValueComission, FindComissionValue *comissionvalue.FindAllByCompanyId) *ComissionValueController {
	return &ComissionValueController{
		CreateComissionValue: createComissionValue,
		FindComissionValue: FindComissionValue,
	}
}
