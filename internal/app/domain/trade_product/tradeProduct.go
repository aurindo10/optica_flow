package tradeproduct

import (
	"time"

	"github.com/google/uuid"
)

type TradeProduct struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CompanyID    string    `json:"company_id"`
	PointAmmount int32     `json:"point_ammount"`
	ImageUrl     *string   `json:"image_url"`
	WhoCreatedID string    `json:"who_created_id"`
}
type TradeProductParams struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CompanyID    string    `json:"company_id"`
	PointAmmount int32     `json:"point_ammount"`
	ImageUrl     *string   `json:"image_url"`
	WhoCreatedID string    `json:"who_created_id"`
}
type TradeProducToUpdate struct {
	ID           uuid.UUID `json:"id"`
	Name         *string    `json:"name"`
	Description  *string    `json:"description"`
	UpdatedAt    time.Time `json:"updated_at"`
	PointAmmount *int32     `json:"point_ammount"`
	ImageUrl     *string   `json:"image_url"`
}



func NewTradeProduct(p *TradeProductParams) *TradeProduct {
	return &TradeProduct{
		ID: uuid.New(),
		Name: p.Name,
		Description: p.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CompanyID: p.CompanyID,
		PointAmmount: p.PointAmmount,
		ImageUrl: p.ImageUrl,
		WhoCreatedID: p.WhoCreatedID,
	}
}

