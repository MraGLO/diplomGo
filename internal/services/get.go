package services

import "diplomGo/pkg/model"

func (s *Services) GetAllSubject() ([]model.Subject, error) {
	return s.Database.db.SelectAllSubject()
}

func (s *Services) GetSubjectByID(id int) (model.Subject, error) {
	return s.Database.db.SelectSubjectByID(id)
}

func (s *Services) GetAllTeachers() ([]model.Teacher, error) {
	return s.Database.db.SelectAllTeachers()
}

func (s *Services) GetTeacherByID(id int) (model.Teacher, error) {
	return s.Database.db.SelectTeacherByID(id)
}

func (s *Services) GetAllTeacherSubject() ([]model.TeacherSubject, error) {
	return s.Database.db.SelectAllTeacherSubject()
}

func (s *Services) GetTeacherSubjectByID(id int) (model.TeacherSubject, error) {
	return s.Database.db.SelectTeacherSubjectByID(id)
}

func (s *Services) GetAllGroups() ([]model.Group, error) {
	return s.Database.db.SelectAllGroups()
}

func (s *Services) GetGroupByID(id int) (model.Group, error) {
	return s.Database.db.SelectGroupByID(id)
}

func (s *Services) GetAllSpecialization() ([]model.Specialization, error) {
	return s.Database.db.SelectAllSpecializations()
}

func (s *Services) GetSpecializationByID(id int) (model.Specialization, error) {
	return s.Database.db.SelectSpecializationByID(id)
}

func (s *Services) GetAllFlows() ([]model.Flow, error) {
	return s.Database.db.SelectAllFlows()
}

func (s *Services) GetFlowByID(id int) (model.Flow, error) {
	return s.Database.db.SelectFlowByID(id)
}

func (s *Services) GetAllStudyPlans() ([]model.StudyPlan, error) {
	return s.Database.db.SelectAllStudyPlans()
}

func (s *Services) GetStudyPlanByID(id int) (model.StudyPlan, error) {
	return s.Database.db.SelectStudyPlanByID(id)
}

func (s *Services) GetAllRecordStudyPlans() ([]model.RecordStudyPlan, error) {
	return s.Database.db.SelectAllRecordStudyPlans()
}

func (s *Services) GetRecordStudyPlanByID(id int) (model.RecordStudyPlan, error) {
	return s.Database.db.SelectRecordStudyPlanByID(id)
}

func (s *Services) GetAllLoads() ([]model.Load, error) {
	return s.Database.db.SelectAllLoads()
}

func (s *Services) GetLoadByID(id int) (model.Load, error) {
	return s.Database.db.SelectLoadByID(id)
}
