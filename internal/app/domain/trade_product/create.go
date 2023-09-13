package tradeproduct

type CreateTradeProduct struct {
	repository Repository
}

func (c *CreateTradeProduct) Execute(p *TradeProductParams) (*TradeProduct, error) {
	tradeProduct := NewTradeProduct(p)
	result, error := c.repository.Create(tradeProduct)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewCreateCreateTradeProduct(repository Repository) *CreateTradeProduct {
	return &CreateTradeProduct{
		repository: repository,
	}
}