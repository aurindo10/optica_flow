package comission

type CreateComission struct {
	repository Repository
}

func (c *CreateComission) Execute(p *CommissionParams) (*Commission, error) {
	request := NewComission(p)
	result, error := c.repository.Create(request)
	if error != nil {
		return nil, error
	}
	response := &Commission{
		ID: result.ID,
		CompanyID: result.CompanyID,
		WhoCreatedID: result.WhoCreatedID,
		WhoUpdatedID: result.WhoUpdatedID,
		Name: result.Name,
		Description: result.Description,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		OrderID: result.OrderID,
		Value: result.Value,
	}
	return response, nil	
}
func NewCreateComission(repository Repository) *CreateComission {
	return &CreateComission{
		repository: repository,
	}
}