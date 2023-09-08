package clientrepository

import (
	"optica_flow/internal/app/domain/client"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewCLientRepository,
		fx.As(new(client.Repository)),
	),
)