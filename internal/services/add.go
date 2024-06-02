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

func (s *Services) AddTableFile(tableFile model.TableFile) (int, error) {
	return s.Database.db.InsertTableFile(&tableFile)
}

func (s *Services) AddSheet(sheet model.Sheet) (int, error) {
	return s.Database.db.InsertSheet(&sheet)
}
func (s *Services) AddRecord(record model.Record) (int, error) {
	return s.Database.db.InsertRecord(&record)
}

func (s *Services) AddSheetRecords(sheetRecords model.SheetRecords) error {
	return s.Database.db.InsertSheetRecords(&sheetRecords)
}
func (s *Services) AddTableFileSheets(tableFileSheet model.TableFileSheet) error {
	return s.Database.db.InsertTableFileSheets(&tableFileSheet)
}
