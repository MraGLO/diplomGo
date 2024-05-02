package model

type Group struct {
	ID               int    `json:"id"`
	SpecializationID int    `json:"specializationid"`
	Name             string `json:"name"`
	Course           int    `json:"course"`
	FlowID           int    `json:"flowId"`
}
