package model

type TeacherPricingRecord struct {
	ID              int `json:"ID"`
	TeacherID       int `json:"teacherID"`
	PricingRecordID int `json:"pricingRecordID"`
}
