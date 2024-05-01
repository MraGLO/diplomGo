package services

import "diplomGo/pkg/model"

func (s *Services) UpdateSubject(id int, subject model.Subject) error {
	return s.Database.db.UpdateSubject(id, &subject)
}

func (s *Services) UpdateTeacher(id int, teacher model.Teacher) error {
	return s.Database.db.UpdateTeacher(id, teacher)
}

func (s *Services) UpdateTeacherSubject(id int, TeacherSubject model.TeacherSubject) error {
	return s.Database.db.UpdateTeacherSubject(id, TeacherSubject)
}
