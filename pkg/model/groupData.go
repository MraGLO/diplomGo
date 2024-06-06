package model

type GroupData struct {
	Group   Group    `json:"group"`
	Records []Record `json:"records"`
}
