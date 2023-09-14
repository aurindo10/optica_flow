package comission


type Repository interface {
	Create(*Commission) (*Commission, error)
	FindByUserId(id string) ([]*Commission, error)
}
