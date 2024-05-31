package model

type Pricing struct {
	ID              int    `json:"id"`
	PricingInfoID   string `json:"pricinginfoid"`
	PricingRecordID int    `json:"pricingrecordid"`
}
