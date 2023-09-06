package repository

import (
	"optica_flow/internal/app/domain/product"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewProductRepository,
		fx.As(new(product.Repository)),
	),
)