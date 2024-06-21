package model

type PricingTableTeacher struct {
	ID             int `json:"ID"`
	PricingTableID int `json:"pricingTableID"`
	TeacherID      int `json:"teacherID"`
}
