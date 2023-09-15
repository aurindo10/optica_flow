package comissionvalue


type CreateValueComission struct {
	repository Repository
}

func (c *CreateValueComission) Execute(r *ComissionValuesParams) (*ComissionValues, error) {
	request := NewComissionValues(r)
	result, error := c.repository.Create(request)
	if error != nil {
		return nil, error
	}
	return result, nil	
}

func NewCreateValueComission(repository Repository) *CreateValueComission {
	return &CreateValueComission{
		repository: repository,
	}
}