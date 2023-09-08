package product

import "go.uber.org/fx"

var Module = fx.Provide(
	NewProduct,
	NewCreateProduct,
	NewGetAllProducts,
	NewUpdateProduct,
	NewDeleteProduct,
	NewGetProduct,
)