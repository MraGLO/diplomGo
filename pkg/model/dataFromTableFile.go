package model

type DataFromTableFile struct {
	TableFile      TableFile       `json:"tablefile"`
	DataFromSheets []DataFromSheet `json:"datafromsheets"`
}
