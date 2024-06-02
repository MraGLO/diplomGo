package model

type Record struct {
	ID              int    `json:"id"`
	SubjectID       int    `json:"subjectid"`
	TimeSemesterOwn string `json:"timesemesterown"`
	TimeSemesterTwo string `json:"timesemestertwo"`
	TeacherID       int    `json:"teacherid"`
}
