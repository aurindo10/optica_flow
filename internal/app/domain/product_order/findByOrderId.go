package productorder

type FindByOrderId struct {
	repository Repository
}

func (c *FindByOrderId) Execute(orderId *string) ([]*ProductOrder, error) {
	result, err := c.repository.FindByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func NewFindByOrderId(repository Repository) *FindByOrderId {
	return &FindByOrderId{repository: repository}
}