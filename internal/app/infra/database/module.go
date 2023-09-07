package database

import (
	repository "optica_flow/internal/app/infra/database/productRepository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewPostgresDatabase,
	),
	repository.Module,)