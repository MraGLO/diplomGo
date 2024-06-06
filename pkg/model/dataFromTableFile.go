package model

type DataFromTableFile struct {
	TableFile TableFile   `json:"tablefile"`
	GroupData []GroupData `json:"groupdata"`
}
