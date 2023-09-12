package points

type FindBySellerId struct {
	repository Repository
}

func (c *FindBySellerId) Execute(id string) ([]*Points, error) {
	result, error := c.repository.FindBySellerId(id)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewFindSellerById(repository Repository) *FindBySellerId {
	return &FindBySellerId{
		repository: repository,
	}
}