package model

type Teacher struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Category   int    `json:"category"`
}
