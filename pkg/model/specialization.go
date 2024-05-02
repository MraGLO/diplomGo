package model

type Specialization struct {
	ID                 int    `json:"id"`
	SpecializationName string `json:"specializationName"`
	QualificationName  string `json:"qualificationName"`
}
