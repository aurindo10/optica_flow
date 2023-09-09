package orders

type CreateOrder struct {
	repository Repository
}

func (c *CreateOrder) Execute(params *Params) (*Order, error) {
	order := NewOrder(params)
	result, error := c.repository.Create(order)
	if error != nil {
		return result, error
	}
	return result, error
}
func NewCreateOrder(repository Repository) *CreateOrder {
	return &CreateOrder{
		repository: repository,
	}
}