package cashflowout

type UpdateCashFlowBalance struct {
	repository Repository
}
func (c *UpdateCashFlowBalance) Execute(p *CashFlowBalanceUpdate) (*CashFlowBalance, error){
	result, error := c.repository.Update(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateCashFlowBalance(repository Repository) *UpdateCashFlowBalance {
	return &UpdateCashFlowBalance{
		repository: repository,
	}
}