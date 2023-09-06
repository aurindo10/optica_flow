package database

import (
	"optica_flow/internal/app/infra/database/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewPostgresDatabase,
	),
	repository.Module,)