package points

type Repository interface {
	Create(*Points) (*Points, error)
	FindBySellerId(id string) ([]*Points, error)
}