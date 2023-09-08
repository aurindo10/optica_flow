package client

type Create struct {
	repository Repository

}



func (c *Create) Execute(client *Params) (*Client,  error) {
	newCliente := NewClient(client)
	result, error := c.repository.Create(newCliente)
	if error != nil {
		return nil, error
	}
	return result, nil
}	
func NewCreate(repository Repository) *Create {
	return &Create{
		repository: repository,
	}
}