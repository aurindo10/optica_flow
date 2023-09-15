package comissionvalue


type Repository interface {
	Create(*ComissionValues) (*ComissionValues, error)
	FindAll(id string) ([]*ComissionValues, error)
}