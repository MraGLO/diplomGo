package services

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
