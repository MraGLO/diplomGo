package model

type TeacherSubject struct {
	ID        int `json:"ID"`
	TeacherID int `json:"teacherID"`
	SubjectID int `json:"subjectID"`
}
