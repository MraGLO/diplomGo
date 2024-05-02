package services

import "diplomGo/pkg/model"

func (s *Services) AddSubject(subject model.Subject) error {
	return s.Database.db.InsertSubject(&subject)
}

func (s *Services) AddTeacher(teacher model.Teacher) error {
	return s.Database.db.InsertTeacher(&teacher)
}

func (s *Services) AddTeacherSubject(teacherSubject model.TeacherSubject) error {
	return s.Database.db.InsertTeacherSubject(&teacherSubject)
}

func (s *Services) AddGroup(group model.Group) error {
	return s.Database.db.InsertGroup(&group)
}

func (s *Services) AddSpecialization(specialization model.Specialization) error {
	return s.Database.db.InsertSpecialization(&specialization)
}

func (s *Services) AddFlow(flow *model.Flow) error {
	return s.Database.db.InsertFlow(flow)
}

func (s *Services) AddStudyPlan(studyPlan *model.StudyPlan) error {
	return s.Database.db.InsertStudyPlan(studyPlan)
}
func (s *Services) AddRecordStudyPlan(recordStudyPlan *model.RecordStudyPlan) error {
	return s.Database.db.InsertRecordStudyPlan(recordStudyPlan)
}

func (s *Services) AddLoad(load *model.Load) error {
	return s.Database.db.InsertLoad(load)
}
