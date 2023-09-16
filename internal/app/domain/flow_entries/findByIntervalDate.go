package flowentries

type FindByIntervalDate struct {
	repository Repository
}

func (c *FindByIntervalDate) Execute(p *FindByDateParams) ([]*CashFlowEntries, error){
	result, error := c.repository.FindByIntervalDate(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewFindByIntervalDate(repository Repository) *FindByIntervalDate {
	return &FindByIntervalDate{
		repository: repository,
	}
}