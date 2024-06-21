package model

type TableFile struct {
	ID   int         `json:"ID"`
	Name string      `json:"name"`
	Date interface{} `json:"date"`
}
