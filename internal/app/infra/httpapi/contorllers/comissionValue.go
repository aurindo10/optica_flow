package controllers

import (
	"errors"
	comissionvalue "optica_flow/internal/app/domain/comission_value"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type ComissionValueController struct {
	CreateComissionValue *comissionvalue.CreateValueComission
	FindComissionValue *comissionvalue.FindAllByCompanyId
	DeleteComissionValue *comissionvalue.DeleteComissionValue
	UpdateComissionValue *comissionvalue.UpdateComissionValue
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
func (r *ComissionValueController) DeleteComissionValueById(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	if error := r.DeleteComissionValue.Execute(idParsed); error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"deleteado com sucesso"})
}
func (r *ComissionValueController) UpdateValueComission(c *fiber.Ctx) error{
	var request comissionvalue.ComissionValuesUpdate
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": error.Error()}) 
	}
	result, error := r.UpdateComissionValue.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": error.Error()}) 
	}
	return c.Status(fiber.StatusOK).JSON(result) 
}
func NewComissionValueController(createComissionValue *comissionvalue.CreateValueComission,
	FindComissionValue *comissionvalue.FindAllByCompanyId,
	deleteComissionValue *comissionvalue.DeleteComissionValue,
	UpdateComissionValue *comissionvalue.UpdateComissionValue,
	) *ComissionValueController {
	return &ComissionValueController{
		CreateComissionValue: createComissionValue,
		FindComissionValue: FindComissionValue,
		DeleteComissionValue: deleteComissionValue,
		UpdateComissionValue: UpdateComissionValue,
	}
}
