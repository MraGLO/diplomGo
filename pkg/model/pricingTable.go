package model

type PricingTable struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Date interface{} `json:"date"`
}
