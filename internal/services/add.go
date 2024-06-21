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

func (s *Services) AddTableFile(tableFile model.TableFile) (int, error) {
	return s.Database.db.InsertTableFile(&tableFile)
}

func (s *Services) AddRecord(record model.Record) (int, error) {
	return s.Database.db.InsertRecord(&record)
}

func (s *Services) AddGroupRecords(groupRecords model.GroupRecords) error {
	return s.Database.db.InsertGroupRecords(&groupRecords)
}

func (s *Services) AddPricingTable(pricingTable model.PricingTable) (int, error) {
	return s.Database.db.InsertPricingTable(&pricingTable)
}

func (s *Services) AddPricingRecord(pricingRecord model.PricingRecord) (int, error) {
	return s.Database.db.InsertPricingRecord(&pricingRecord)
}
func (s *Services) AddTeacherPricingRecord(teacherPricingRecord model.TeacherPricingRecord) error {
	return s.Database.db.InsertTeacherPricingRecord(&teacherPricingRecord)
}
