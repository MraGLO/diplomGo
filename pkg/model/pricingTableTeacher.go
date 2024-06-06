package model

type PricingTableTeacher struct {
	ID             int `json:"id"`
	PricingTableID int `json:"pricingtableid"`
	TeacherID      int `json:"teacherid"`
}
