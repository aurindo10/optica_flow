package controllers

import (
	"optica_flow/internal/app/domain/points"

	"github.com/gofiber/fiber/v2"
)

type PointsController struct {
	CreatPoint *points.CreatePoints
	FindBySellerId *points.FindBySellerId
}

func (r *PointsController) CreatePoints(c *fiber.Ctx) error {
	var request points.PointsParams
	error := c.BodyParser(&request)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error ao obter informações"})
	}
	err := r.Validate(&request)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	result, err := r.CreatPoint.Execute(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}
func (r *PointsController) FindBySeller(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(id) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "id é obrigatório"})
	}
	result, error := r.FindBySellerId.Execute(id)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "houve um erro ao tentar obter os pontos"})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func NewPointsController(CreatPoint *points.CreatePoints, FindBySellerId *points.FindBySellerId) *PointsController {
	return &PointsController{
		CreatPoint: CreatPoint,
		FindBySellerId: FindBySellerId,
	}
}

func (r *PointsController) Validate(p *points.PointsParams) error {
	if p.Ammount == 0 {
		return fiber.NewError(400, "Quantidade de pontos é obrigatório")
	}
	if p.CompanyID == "" {
		return fiber.NewError(400, "id da empresa é obrigatorio")
	}
	if p.Name == "" {
		return fiber.NewError(400, "nome dos pontos é obrigatorio")
	}
	if p.OrderID == "" {
		return fiber.NewError(400, "id do order é obrigatório")
	}
	if p.ValidDate.IsZero() {
		return fiber.NewError(400, "Data de validade é obrigatório")
	}
	if p.SellerID == "" {
		return fiber.NewError(400, "id do usuário é obrigatório")
	}
	return nil
}