package controllers

import (
	"optica_flow/internal/app/domain/fornecedor"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type FornecedorController struct {
	createFornecedor *fornecedor.CreateFornecedor
	getFornecedorById *fornecedor.GetFornecedorById
	findAllFornecedores *fornecedor.Find
	update *fornecedor.Update
	delete *fornecedor.Delete
}


func (p *FornecedorController) CreateFornecedor(c *fiber.Ctx)  error {
	var body fornecedor.Fornecedor
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}
	if body.Name == "" || body.Telefone == "" || body.Email == "" || body.Adress == "" ||
		body.CompanyID == "" || body.WhoCreatedID == "" || body.WhoUpdatedID == "" || body.Cnpj == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	}
	response, error := p.createFornecedor.Execute(&body)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	return  c.Status(fiber.StatusCreated).JSON(response)
}
func (p *FornecedorController) GetFornecedorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == uuid.Nil.String() {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	}
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	response, error := p.getFornecedorById.Execute(idParsed)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	return  c.Status(fiber.StatusOK).JSON(response)
}

func (p *FornecedorController) FindAllFornecedoresByCompanyId(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	}
	response, error := p.findAllFornecedores.Execute(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	return  c.Status(fiber.StatusOK).JSON(response)
}

func (p *FornecedorController) UpdateFornecedor (c *fiber.Ctx) error {
	var body fornecedor.FornecedorToUpdate
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}
	if body.ID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	}
	response, error := p.update.Execute(&body)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	return  c.Status(fiber.StatusOK).JSON(response)
}
func (p *FornecedorController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == uuid.Nil.String() {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	}
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	error = p.delete.Execute(idParsed)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": error.Error()})
	}
	return  c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "deleted"})
}
func NewFornecedorController(createFornecedor *fornecedor.CreateFornecedor, getFornecedorById *fornecedor.GetFornecedorById, findAllFornecedores *fornecedor.Find, update *fornecedor.Update, delete *fornecedor.Delete) *FornecedorController{
	return &FornecedorController{
		createFornecedor: createFornecedor,
		getFornecedorById: getFornecedorById,
		findAllFornecedores: findAllFornecedores,
		update: update,
		delete: delete,
	}
}