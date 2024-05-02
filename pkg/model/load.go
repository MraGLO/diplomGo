package model

type Load struct {
	ID                int `json:"id"`
	RecordStudyPlanID int `json:"recordStudyPlanId"`
	TimeOnSubject     int `json:"timeOnSubject"`
	TeacherID         int `json:"teacherId"`
	GroupID           int `json:"groupId"`
}
