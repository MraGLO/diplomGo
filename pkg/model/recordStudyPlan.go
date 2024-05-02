package model

type RecordStudyPlan struct {
	ID                        int `json:"id"`
	Semester                  int `json:"semester"`
	SubjectID                 int `json:"subjectId"`
	MaxTime                   int `json:"maxTime"`
	Independet                int `json:"independet"`
	Consultations             int `json:"consultations"`
	RequiredTime              int `json:"requiredTime"`
	LecturesAndLessons        int `json:"lecturesAndLessons"`
	PracticalExercises        int `json:"practicalExercises"`
	LaboratoryExercises       int `json:"laboratoryExercises"`
	Seminars                  int `json:"seminars"`
	CourseProject             int `json:"courseProject"`
	IntermediateCertification int `json:"intermediateCertification"`
	StudyPlanID               int `json:"studyPlanID"`
}
