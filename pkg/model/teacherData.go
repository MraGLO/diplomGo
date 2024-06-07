package model

type TeacherData struct {
	Teacher          Teacher            `json:"teacher"`
	RecordForPricing []RecordForPricing `json:"recordforpricing"`
}
