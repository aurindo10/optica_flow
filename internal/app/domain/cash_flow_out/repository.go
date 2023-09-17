package cashflowout

type Repository interface {
	Create(*CashFlowBalance) (*CashFlowBalance, error)
}