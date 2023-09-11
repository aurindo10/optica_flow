package productorder

import "go.uber.org/fx"

var Module = fx.Provide(
	NewProductOrder,
	NewCreateProductOrder,
	NewFindByOrderId,
	NewUpdateProductOrder,
	NewDeleteProductOrder,
)

