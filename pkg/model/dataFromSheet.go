package model

type DataFromSheet struct {
	Sheet   Sheet    `json:"sheet"`
	Records []Record `json:"records"`
}
