package model

type DataFromPricingTable struct {
	PricingTable PricingTable  `json:"pricingtable"`
	TeacherData  []TeacherData `json:"teacherdata"`
}
