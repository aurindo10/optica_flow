package fornecedor

import "go.uber.org/fx"

var Module = fx.Provide(
	NewFornecedor,
	NewCreateFornecedor,
)