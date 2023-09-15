package comissionvalue


type FindAllByCompanyId struct {
	repository Repository
}

func (c *FindAllByCompanyId) Execute(id string) ([]*ComissionValues, error) {
	result, error := c.repository.FindAll(id)
	if error != nil {
		return nil, error
	}
	return result, nil	
}

func NewFindAllByCompanyId(repository Repository) *FindAllByCompanyId {
	return &FindAllByCompanyId{
		repository: repository,
	}
}