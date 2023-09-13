package comission


type Repository interface {
	Create(*Commission) (*Commission, error)
}
