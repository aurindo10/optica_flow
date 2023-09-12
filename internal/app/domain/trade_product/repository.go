package tradeproduct

type Repository interface {
	Create(*TradeProduct) (*TradeProduct, error)
	FindAll(id string) ([]*TradeProduct, error)
}