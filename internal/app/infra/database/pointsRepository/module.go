package pointsrepository

import (
	"optica_flow/internal/app/domain/points"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPointsRepository,
		fx.As(new(points.Repository)),
	),
)