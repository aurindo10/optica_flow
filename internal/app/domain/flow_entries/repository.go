package flowentries

type Repository interface {
	Create(*CashFlowEntries) (*CashFlowEntries, error)
}

