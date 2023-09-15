package comissionvalue


type Repository interface {
	Create(*ComissionValues) (*ComissionValues, error)
}