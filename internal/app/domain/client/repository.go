package client

import "github.com/google/uuid"

type Repository interface {
	Create(client *Client) (*Client,error)
	Find(id string) ([]*Client, error)
	Update(client *ClientToUpdate) (*Client, error)
	FindOne (id uuid.UUID) (*Client, error)
}

