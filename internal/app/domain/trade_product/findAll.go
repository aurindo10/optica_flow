package tradeproduct

type FindAllRepository struct {
	repository Repository
}

func (c *FindAllRepository) Execute(id string) ([]*TradeProduct, error) {
	result, error := c.repository.FindAll(id)
	if error != nil {
		return nil, error
	}
	return result, nil
}
func NewFindAllRepository(repository Repository) *FindAllRepository {
	return &FindAllRepository{
		repository: repository,
	}
}