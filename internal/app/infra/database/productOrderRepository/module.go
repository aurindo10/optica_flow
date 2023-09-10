package productorderrepository

import (
	productorder "optica_flow/internal/app/domain/product_order"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewProductOrderRepository,
		fx.As(new(productorder.Repository)),
	),
)