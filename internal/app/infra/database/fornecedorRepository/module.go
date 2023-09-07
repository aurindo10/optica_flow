package fornecedorrepository

import (
	"optica_flow/internal/app/domain/fornecedor"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewFornecedorRepository,
		fx.As(new(fornecedor.Repository)),
	),
)