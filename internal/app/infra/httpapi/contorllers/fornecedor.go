package controllers

import (
	"optica_flow/internal/app/domain/fornecedor"

	"github.com/gofiber/fiber/v2"
)


type FornecedorController struct {
	createFornecedor *fornecedor.CreateFornecedor
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

func NewFornecedorController(createFornecedor *fornecedor.CreateFornecedor) *FornecedorController{
	return &FornecedorController{
		createFornecedor: createFornecedor,
	}
}