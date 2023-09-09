package controllers

import (
	"fmt"
	"optica_flow/internal/app/domain/client"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type ClientController struct {
	createClient *client.Create
	findClients *client.FindClients
	updateClient *client.UpdateClient
	findOneClient *client.FindOne
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
func (r *ClientController) Update(c *fiber.Ctx) error {
	var request client.ClientToUpdate
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	if request.ID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "id field is required"})
	}
	response, err := r.updateClient.Execute(&request)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
func (r *ClientController) FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON("id não pode ser vazio")
	}
	idParsed, error := uuid.Parse(id)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	result, error := r.findOneClient.Execute(idParsed)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func NewClientController(createClient *client.Create, findClients *client.FindClients, updateClient *client.UpdateClient, findOneClient *client.FindOne) *ClientController {
	return &ClientController{
		createClient: createClient,
		findClients: findClients,
		updateClient: updateClient,
		findOneClient: findOneClient,
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