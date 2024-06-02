package model

type TableFile struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	Date interface{} `json:"date"`
}
