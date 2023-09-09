package orders


type Repository interface {
	Create (*Order) (*Order, error)
}