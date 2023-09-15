package comissionvaluerepository

import (
	comissionvalue "optica_flow/internal/app/domain/comission_value"

	"go.uber.org/fx"
)


var Module = fx.Provide(
	fx.Annotate(
		NewComissionValueRepository,
		fx.As(new(comissionvalue.Repository)),
	),
)