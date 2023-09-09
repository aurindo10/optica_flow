package controllers

import (
	"fmt"
	"optica_flow/internal/app/domain/client"

	"github.com/gofiber/fiber/v2"
)
type ClientController struct {
	createClient *client.Create
	findClients *client.FindClients
}
func (r *ClientController) Create(c *fiber.Ctx) error {
	var request client.Params
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	error := Validade(&request)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	response, err := r.createClient.Execute(&request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r *ClientController) Find(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON("id não pode ser vazio")
	}
	result, error := r.findClients.FindAll(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func NewClientController(createClient *client.Create, findClients *client.FindClients) *ClientController {
	return &ClientController{
		createClient: createClient,
		findClients: findClients,
	}
}

func Validade(p *client.Params) error {
	if p.FullName == "" {
		return fmt.Errorf("FullName não pode ser vazio")
	}
	if p.Telefone == "" {
		return fmt.Errorf("telefone não pode ser vazio")
	}
	if p.Email == "" {
		return fmt.Errorf("email não pode ser vazio")
	}
	if p.BirthDate.IsZero() {
		return fmt.Errorf("eirthDate não pode ser vazio")
	}
	if p.CompanyID == "" {
		return fmt.Errorf("companyID não pode ser vazio")
	}
	if p.WhoCreatedID == "" {
		return fmt.Errorf("hhoCreatedID não pode ser vazio")
	}
	if p.WhoUpdatedID == "" {
		return fmt.Errorf("whoUpdatedID não pode ser vazio")
	}
	return nil
}