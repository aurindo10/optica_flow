package cashflowout

type FindByRangeDate struct {
	repository Repository
}

func (c *FindByRangeDate) Execute(p *FindByRangeDateParams) ([]*CashFlowBalance, error) {
	result, error := c.repository.FindByRangeDate(p)
	if error != nil {
		return nil, error 
	}
	return result, error
}

func NewFindByRangeDate(repository Repository) *FindByRangeDate {
	return &FindByRangeDate{
		repository: repository,
	}
}