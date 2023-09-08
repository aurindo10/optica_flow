package controllers

import (
	"optica_flow/internal/app/domain/fornecedor"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type FornecedorController struct {
	createFornecedor *fornecedor.CreateFornecedor
	getFornecedorById *fornecedor.GetFornecedorById
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

func NewFornecedorController(createFornecedor *fornecedor.CreateFornecedor, getFornecedorById *fornecedor.GetFornecedorById) *FornecedorController{
	return &FornecedorController{
		createFornecedor: createFornecedor,
		getFornecedorById: getFornecedorById,
	}
}