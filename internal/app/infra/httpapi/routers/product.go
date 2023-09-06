package routers

import (
	controllers "optica_flow/internal/app/infra/httpapi/contorllers"
	"optica_flow/internal/app/infra/httpapi/middleware"

	"github.com/gofiber/fiber/v2"
)


type ProductRouter struct {
	controller *controllers.ProductController
	authMiddleware *middleware.AuthMiddleware
}


func (p *ProductRouter) Load(r *fiber.App) {
	r.Post("/product",p.controller.CreateProduct)
	r.Get("/products", p.controller.GetAllProducts)
	r.Put("/product", p.controller.UpdateProduct)
	r.Delete("/product/:id", p.controller.DeleteProduct)
}

func NewProductRouter(
	controller *controllers.ProductController, authMiddleware *middleware.AuthMiddleware,
) *ProductRouter {
	return &ProductRouter{
		controller: controller,
		authMiddleware: authMiddleware,
	}
}