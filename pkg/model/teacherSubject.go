package model

type TeacherSubject struct {
	ID        int `json:"id"`
	TeacherID int `json:"teacherid"`
	SubjectID int `json:"subjectid"`
}
