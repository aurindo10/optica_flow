package client

type Repository interface {
	Create(client *Client) (*Client,error)
}

