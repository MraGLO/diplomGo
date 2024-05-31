package model

type GroupTableOfSubject struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"groupid"`
	Name       string `json:"name"`
	CreateDate int    `json:"createdate"`
}
