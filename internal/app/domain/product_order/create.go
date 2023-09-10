package productorder

type CreateProductOrder struct {
	repository Repository
}

func (c *CreateProductOrder) Execute(p *ProductOrderParms) (*ProductOrder, error) {
	product := NewProductOrder(p)
	result, err := c.repository.Create(product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewCreateProductOrder(repository Repository) *CreateProductOrder {
	return &CreateProductOrder{
		repository: repository,
	}
}