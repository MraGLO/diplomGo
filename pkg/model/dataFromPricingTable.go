package model

type DataFromPricingTable struct {
	PricingTable PricingTable  `json:"pricingTable"`
	TeacherData  []TeacherData `json:"teacherData"`
}
