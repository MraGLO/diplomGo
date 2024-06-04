package model

type NameHZ struct {
	ID          int    `json:"id"`
	TableFileID int    `json:"specializationid"`
	Name        string `json:"name"`
	Course      int    `json:"course"`
}
