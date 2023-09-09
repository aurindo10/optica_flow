package orders


type UpdateOrder struct {
	repository Repository
}

func (c *UpdateOrder) Execute(p *OrderToUpdate) (*Order, error) {
	order, error := c.repository.Update(p)
	if error != nil {
		return nil, error
	}
	return order, nil
}

func NewUpdateOrder(repository Repository) *UpdateOrder {
	return &UpdateOrder{
		repository: repository,
	}
}	