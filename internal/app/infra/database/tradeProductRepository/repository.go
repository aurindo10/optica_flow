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

func NewTradeProductRepository(db *database.Queries) *TradeProductRepository {
	return &TradeProductRepository{
		db: db,
	}
}