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
	response := &TradeProduct{
		ID: result.ID,
		Name: result.Name,
		Description: result.Description,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		CompanyID: result.CompanyID,
		PointAmmount: result.PointAmmount,
		ImageUrl: result.ImageUrl,
		WhoCreatedID: result.WhoCreatedID,
	}
	return response, nil
}

func NewCreateCreateTradeProduct(repository Repository) *CreateTradeProduct {
	return &CreateTradeProduct{
		repository: repository,
	}
}