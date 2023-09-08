package client

type Create struct {
	repository Repository
}
func (c *Create) Execute(client *Client) (*Client,  error) {
	result, error := c.repository.Create(client)
	if error != nil {
		return nil, error
	}
	return &result, nil
}	
func NewCreate(repository Repository) *Create {
	return &Create{
		repository: repository,
	}
}