package flowentries

type CreateFlowEntrie struct {
	repository Repository
}

func (c *CreateFlowEntrie) Execute(p *CashFlowEntriesParams) (*CashFlowEntries, error) {
	request := NewCashFlowEntries(p)
	result, error := c.repository.Create(request)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewCreateFlowEntrie(repository Repository) *CreateFlowEntrie {
	return &CreateFlowEntrie{
		repository: repository,
	}
}