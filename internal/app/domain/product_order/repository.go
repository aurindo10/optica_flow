package productorder

type Repository interface {
	Create(p *ProductOrder) (*ProductOrder, error)
}
