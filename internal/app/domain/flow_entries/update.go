package flowentries


type UpdateFlowEntrie struct {
	repository Repository
}

func (c *UpdateFlowEntrie) Execute(p *CashFlowEntriesUpdate) (*CashFlowEntries, error) {
	result, error := c.repository.Update(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateFlowEntrie(repository Repository) *UpdateFlowEntrie {
	return &UpdateFlowEntrie{
		repository: repository,
	}
}