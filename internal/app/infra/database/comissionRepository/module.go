package comissionrepository

import (
	"optica_flow/internal/app/domain/comission"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewComissionRepository,
		fx.As(new(comission.Repository)),
	),
)