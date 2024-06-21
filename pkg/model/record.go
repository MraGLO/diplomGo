package model

type Record struct {
	ID              int    `json:"ID"`
	SubjectID       int    `json:"subjectID"`
	TimeSemesterOwn string `json:"timeSemesterOwn"`
	TimeSemesterTwo string `json:"timeSemesterTwo"`
	TeacherID       int    `json:"teacherID"`
	TableFileID     int    `json:"tableFileID"`
}
