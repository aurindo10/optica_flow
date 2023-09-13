package tradeproduct

type UpdateTradeProduct struct {
	repository Repository
}

func(c *UpdateTradeProduct) Execute(p *TradeProducToUpdate) (*TradeProduct, error){
	result, error := c.repository.Update(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateTradeProduct(repository Repository) *UpdateTradeProduct {
	return &UpdateTradeProduct{
		repository: repository,
	}
}