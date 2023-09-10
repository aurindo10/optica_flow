package orders

type FindAllOrders struct {
	repository Repository
}

func (c *FindAllOrders) Execute(companyId string) ([]*Order, error) {
	result, error := c.repository.FindAll(companyId)
	if error != nil {
		return nil, error
	}
	return result, nil
}
func NewFindAllOrder(repository Repository) *FindAllOrders{
	return &FindAllOrders{repository: repository}
}