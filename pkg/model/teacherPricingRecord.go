package model

type TeacherPricingRecord struct {
	ID              int `json:"id"`
	TeacherID       int `json:"teacherid"`
	PricingRecordID int `json:"pricingrecordid"`
}
