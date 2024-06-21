package model

type DataFromTableFile struct {
	TableFile TableFile   `json:"tableFile"`
	GroupData []GroupData `json:"groupData"`
}
