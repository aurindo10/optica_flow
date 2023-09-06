package product

type UpdateProduct struct {
	repository Repository
}

func (c *UpdateProduct) Execute(request *ProductRequestToUpdate)  (*Product, error) {
	product, err := c.repository.UpdateProduct(request)
	if err != nil {
		return  nil, err
	}
	return  product, nil
}
func NewUpdateProduct(repository Repository) *UpdateProduct {
	return &UpdateProduct{
		repository: repository,
	}
}