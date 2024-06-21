package model

type PricingTable struct {
	ID   int         `json:"ID"`
	Name string      `json:"name"`
	Date interface{} `json:"date"`
}
