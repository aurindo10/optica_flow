package controllers

import (
	"errors"
	"optica_flow/internal/app/domain/comission"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type ComissionController struct {
	Create *comission.CreateComission
	FindByUserId *comission.FindByUserId
	DeleteById *comission.DeleteComission
	Update *comission.UpdateComission
}

func (r *ComissionController) CreateComission(c *fiber.Ctx) error {
	var request comission.CommissionParams
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "error ao tentar obter o body"})
	}
	if error := r.Validate(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	result, error := r.Create.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (r *ComissionController) FindByUser(c *fiber.Ctx) error {
	id := c.Params("id")
	result, error := r.FindByUserId.Execute(id)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (r *ComissionController) DeleteComissionById(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}

	if error := r.DeleteById.Execute(idParsed); error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error.Error()) 
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"deleteado com sucesso"})
}

func (r *ComissionController) UpdateComissionById(c *fiber.Ctx) error {
	var request comission.CommissionToUpdate
	if error := c.BodyParser(&request); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	result, error := r.Update.Execute(&request)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func NewComissionCOntroller(Create *comission.CreateComission, findByUserId *comission.FindByUserId,
	 deleteById *comission.DeleteComission, update *comission.UpdateComission) *ComissionController {
	return &ComissionController{
		Create: Create,
		FindByUserId: findByUserId,
		DeleteById: deleteById,
		Update: update,
	}
}

func (r *ComissionController) Validate(request *comission.CommissionParams) error {
	if request.Name == "" {
		return errors.New("nome é obrigatório") 
	}
	if request.Description == "" {
		return errors.New("descrição é obrigatória")
	}
	if request.CompanyID == "" {
		return errors.New("companyID é obrigatório")
	}
	if request.WhoCreatedID == "" {
		return errors.New("whoCreatedID é obrigatório")
	}
	if request.WhoUpdatedID == "" {
		return errors.New("whoUpdatedID é obrigatório")
	}
	if request.Value == 0.0 {
		return errors.New("value não pode ser zero")
	}
	return nil
}
