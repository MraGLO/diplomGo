package model

type TableOfSubjectsRecord struct {
	ID              int `json:"id"`
	SubjectID       int `json:"subjectid"`
	TimeSemesterOwn int `json:"timesemesterown"`
	TimeSemesterTwo int `json:"timesemestertwo"`
	TeacherID       int `json:"teacherid"`
}
