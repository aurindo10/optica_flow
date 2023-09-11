package productorder

type UpdateProductOrder struct {
	repository Repository
}

func (c *UpdateProductOrder) Execute(p *ProductOrderToUpdate) (*ProductOrder, error) {
	result, error := c.repository.UpdateProductOrder(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateProductOrder(repository Repository) *UpdateProductOrder {
	return &UpdateProductOrder{repository: repository}
}