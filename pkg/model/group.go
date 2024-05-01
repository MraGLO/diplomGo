package model

type Group struct {
	ID               int    `json:"id"`
	SpecializationID int    `json:"specializationid"`
	Name             string `json:"name"`
}
