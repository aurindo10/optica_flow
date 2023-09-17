package cashflowout


type CreateFlowBalance struct {
	repository Repository
}

func (c *CreateFlowBalance) Execute(p *CashFlowBalanceParams) (*CashFlowBalance, error) {
	request := NewCashFlowBalance(p)
	result, error := c.repository.Create(request)
	if error != nil {
		return nil, error
	}
	return result, error
}


func NewCreateFlowBalance(repository Repository) *CreateFlowBalance {
	return &CreateFlowBalance{
		repository: repository,
	}
}