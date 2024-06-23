package model

type RecordForTableFileView struct {
	GroupID         int    `json:"groupID"`
	GroupName       string `json:"groupName"`
	SubjectID       int    `json:"subjectID"`
	SubjectName     string `json:"subjectName"`
	TimeSemesterOwn string `json:"timeSemesterOwn"`
	TimeSemesterTwo string `json:"timeSemesterTwo"`
}
