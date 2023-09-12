package points

type Repository interface {
	Create(*Points) (*Points, error)
}