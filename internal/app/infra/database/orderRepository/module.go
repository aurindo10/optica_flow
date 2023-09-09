package orderrepository

import (
	"optica_flow/internal/app/domain/orders"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewOrderRepository,
		fx.As(new(orders.Repository)),
	),
)