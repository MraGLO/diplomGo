package model

type RecordForPricing struct {
	GroupName       string `json:"groupName"`
	SubjectName     string `json:"subjectName"`
	TimeSemesterOwn string `json:"timeSemesterOwn"`
	TimeSemesterTwo string `json:"timeSemesterTwo"`
}
