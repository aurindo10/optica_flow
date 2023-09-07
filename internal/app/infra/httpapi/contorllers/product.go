package controllers

import (
	"optica_flow/internal/app/domain/product"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	createProduct *product.CreateProduct
	getAllProducts *product.GetAllProducts
	updateProduct *product.UpdateProduct
	deleteProduct *product.DeleteProduct
}
type CreateProductRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	FornecedorID *string   `json:"fornecedor_id"`
	Description  string    `json:"description"`
	Brand        string    `json:"brand"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	BarCode      string    `json:"bar_code"`
	Quantity     int32     `json:"quantity"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
}

func (p *ProductController) CreateProduct(c *fiber.Ctx) error {
	var request CreateProductRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}
	if request.Name == "" || request.Price == 0 || request.FornecedorID == nil ||
	 	request.Description == "" || request.Brand == "" || request.BarCode == "" ||
		request.Quantity == 0 || request.CompanyID == "" || request.WhoCreatedID == "" || request.WhoUpdatedID =="" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "some fields are required"})
	
	}
	product, error := p.createProduct.Execute(&product.CreateProductParams{
		ID:          request.ID,
		Name:        request.Name,
		Price:       request.Price,
		FornecedorID: request.FornecedorID,
		Description: request.Description,
		Brand:       request.Brand,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		BarCode:     request.BarCode,
		Quantity:    request.Quantity,
		CompanyID:  request.CompanyID,
		WhoCreatedID:request.WhoCreatedID,
		WhoUpdatedID:request.WhoUpdatedID,
	})
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}

func (p *ProductController) GetAllProducts(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "id is required"},
		)
	}
	products, error := p.getAllProducts.Execute(id)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": error.Error()},
		)
	}
	return c.Status(fiber.StatusOK).JSON(products)
}
func (p *ProductController) UpdateProduct(c *fiber.Ctx) error{
	info := &product.ProductRequestToUpdate{}
	if error := c.BodyParser(info); error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": error.Error()},
		)
	}
	product, error := p.updateProduct.Execute(info)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": error.Error()},
		)
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}
func (p * ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed := uuid.MustParse(id)
	if error := p.deleteProduct.Execute(idParsed); error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": error.Error()},
		)
	}
	return nil
}
func NewProductController(createProduct *product.CreateProduct,
	 getAllProducts *product.GetAllProducts, updateProduct *product.UpdateProduct, deleteProduct *product.DeleteProduct) *ProductController {
	return &ProductController{createProduct, getAllProducts, updateProduct, deleteProduct}
}


