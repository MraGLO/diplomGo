package services

import "diplomGo/pkg/model"

func (s *Services) DeleteSubject(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteSubject(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTeacher(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTeacher(id)
	if err != nil || countAffected == 0 {
		return false, err
	}
	return true, nil
}

func (s *Services) DeleteTeacherSubject(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTeacherSubject(id)
	if err != nil || countAffected == 0 {
		return false, err
	}
	return true, nil
}

func (s *Services) DeleteRecord(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteRecord(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTableFile(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTableFile(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteGroupRecords(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteGroupRecords(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeletePricingTable(id int) (bool, error) {
	countAffected, err := s.Database.db.DeletePricingTable(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeletePricingRecord(id int) (bool, error) {
	countAffected, err := s.Database.db.DeletePricingRecord(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTeacherPricingRecord(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTeacherPricingRecord(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTeacherSubjectByData(teacherSubject model.TeacherSubject) (bool, error) {
	countAffected, err := s.Database.db.DeleteTeacherSubjectByData(teacherSubject)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTeacherPricingRecordByTeacherID(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTeacherPricingRecordByTeacherID(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}
