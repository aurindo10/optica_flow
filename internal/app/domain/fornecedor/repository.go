package fornecedor

type Repository interface {
	CreateFornecedor(fornecedor *Fornecedor) (*Fornecedor, error)
}
