package comissionvalue

import "github.com/google/uuid"

type ComissionValues struct {
	ID           uuid.UUID `json:"id"`
	Percentage   float64   `json:"percentage"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	Type         string    `json:"type"`
}
type ComissionValuesParams struct {
	ID           uuid.UUID `json:"id"`
	Percentage   float64   `json:"percentage"`
	CompanyID    string    `json:"company_id"`
	WhoCreatedID string    `json:"who_created_id"`
	WhoUpdatedID string    `json:"who_updated_id"`
	Type         string    `json:"type"`
}
type ComissionValuesUpdate struct {
	ID           uuid.UUID `json:"id"`
	Percentage   *float64   `json:"percentage"`
	CompanyID    *string    `json:"company_id"`
	WhoUpdatedID *string    `json:"who_updated_id"`
	Type         *string    `json:"type"`
}
func NewComissionValues(p *ComissionValuesParams) *ComissionValues {
	return &ComissionValues{
		ID: uuid.New(),
		Percentage: p.Percentage,
		CompanyID: p.CompanyID,
		WhoCreatedID: p.WhoCreatedID,
		WhoUpdatedID: p.WhoUpdatedID,
		Type: p.Type,
	}
}