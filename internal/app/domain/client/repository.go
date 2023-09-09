package client

type Repository interface {
	Create(client *Client) (*Client,error)
	Find(id string) ([]*Client, error)
}

