package model

type RecordForPricing struct {
	GroupName       string `json:"groupname"`
	SubjectName     string `json:"subjectname"`
	TimeSemesterOwn string `json:"timesemesterown"`
	TimeSemesterTwo string `json:"timesemestertwo"`
}
