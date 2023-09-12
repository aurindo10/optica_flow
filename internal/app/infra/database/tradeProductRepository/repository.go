package tradeproductrepository

import (
	"context"
	tradeproduct "optica_flow/internal/app/domain/trade_product"
	database "optica_flow/internal/app/infra/database/queries"
)

type TradeProductRepository struct {
	db *database.Queries
}

func (c *TradeProductRepository) Create(p *tradeproduct.TradeProduct) (*tradeproduct.TradeProduct, error) {
	request := database.CreateTradePdrouctParams{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		CompanyID: p.CompanyID,
		PointAmmount: p.PointAmmount,
		ImageUrl: p.ImageUrl,
		WhoCreatedID: p.WhoCreatedID,
	}
	result, error := c.db.CreateTradePdrouct(context.Background(), request)
	if error != nil {
		return nil, error
	}
	response := &tradeproduct.TradeProduct{
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
func (c *TradeProductRepository) FindAll(id string) ([]*tradeproduct.TradeProduct, error) {
	result, error := c.db.FindAllTradeProducts(context.Background(), id)
	if error != nil {
		return nil, error
	}
	var response []*tradeproduct.TradeProduct
	for _, v := range result {
		response = append(response, &tradeproduct.TradeProduct{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			CompanyID: v.CompanyID,
			PointAmmount: v.PointAmmount,
			ImageUrl: v.ImageUrl,
			WhoCreatedID: v.WhoCreatedID,
		},)
	}
	return response, nil
}

func NewTradeProductRepository(db *database.Queries) *TradeProductRepository {
	return &TradeProductRepository{
		db: db,
	}
}