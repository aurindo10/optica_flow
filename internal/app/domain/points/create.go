package points


type CreatePoints struct {
	repository Repository
}

func (c *CreatePoints) Execute(p *PointsParams) (*Points, error) {
	points := NewPoints(p)
	result, error := c.repository.Create(points)
	if error != nil {
		return nil, error
	}
	return result, error
}

func NewCreatePoints(repository Repository) *CreatePoints {
	return &CreatePoints{repository: repository}
}