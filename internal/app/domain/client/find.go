package client

type FindClients struct {
	repository Repository
}

func (c *FindClients) FindAll(id string) ([]*Client, error) {
	result, error := c.repository.Find(id)
	if error != nil {
		return nil, error
	}
	return result, nil  
}


func NewFindClients(repository Repository) *FindClients {
	return &FindClients{
		repository: repository,
	}
}