package comission

import "time"

type UpdateComission struct {
	repository Repository
}

func (c *UpdateComission) Execute(r *CommissionToUpdate) (*Commission, error) {
	request := CommissionToUpdate{
		ID: r.ID,
		Name: r.Name,
		Description: r.Description,
		WhoUpdatedID: r.WhoUpdatedID,
		Value: r.Value,
		UpdatedAt: time.Now(),
	}
	result, error := c.repository.UpdateById(&request)
	if error != nil {
		return nil, error
	}
	return result, nil
}

func NewUpdateComission(repository Repository) *UpdateComission {
	return &UpdateComission{
		repository: repository,
	}
}