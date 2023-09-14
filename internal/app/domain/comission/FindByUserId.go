package comission


type FindByUserId struct {
	repository Repository
}

func (c *FindByUserId) Execute(id string) ([]*Commission, error) {
	result, error := c.repository.FindByUserId(id)
	if error != nil {
		return nil, error
	}
	return result, error
}

func NewFindByUserId(repository Repository) *FindByUserId {
	return &FindByUserId{
		repository: repository,
	}
}