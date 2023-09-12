package tradeproduct

type Repository interface {
	Create(*TradeProduct) (*TradeProduct, error)
}