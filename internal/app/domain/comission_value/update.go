package comissionvalue

type UpdateComissionValue struct {
	repository Repository
}

func (c *UpdateComissionValue) Execute(p *ComissionValuesUpdate) (*ComissionValues, error) {
	result, error := c.repository.UpdateComissionValue(p)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateComissionValue(repository Repository) *UpdateComissionValue {
	return &UpdateComissionValue{
		repository: repository,
	}
}