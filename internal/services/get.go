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

func (s *Services) GetIDBySubject(subj string) (int, error) {
	return s.Database.db.SelectIDBySubject(subj)
}
func (s *Services) GetIDByTeacherSurname(surname string) (int, error) {
	return s.Database.db.SelectIDByTeacherSurname(surname)
}

func (s *Services) GetIDByGroupName(groupName string) (int, error) {
	return s.Database.db.SelectIDByGroupName(groupName)
}

func (s *Services) GetAllTableFiles() ([]model.TableFile, error) {
	return s.Database.db.SelectAllTableFiles()
}

func (s *Services) GetTableFilesByID(id int) (model.TableFile, error) {
	return s.Database.db.SelectTableFilesByID(id)
}

// func (s *Services) GetAllTableFileGroupsByTableFileID(tableFileID int) ([]model.TableFileGroup, error) {
// 	return s.Database.db.SelectAllTableFileGroupsByTableFileID(tableFileID)
// }

func (s *Services) GetAllGroupRecordsByGroupID(groupID int) ([]model.GroupRecords, error) {
	return s.Database.db.SelectAllGroupRecordsByGroupID(groupID)
}

func (s *Services) GetRecordByID(id int) (model.Record, error) {
	return s.Database.db.SelectRecordByID(id)
}
func (s *Services) GetAllTeachersFromPriceView(tableFileID int) ([]model.Teacher, error) {
	return s.Database.db.SelectAllTeachersFromPriceView(tableFileID)
}
func (s *Services) GetAllRecordForPricingFromPriceView(tableFileID int, teacherID int) ([]model.RecordForPricing, error) {
	return s.Database.db.SelectAllRecordForPricingFromPriceView(tableFileID, teacherID)
}
