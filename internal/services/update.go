package services

import "diplomGo/pkg/model"

func (s *Services) UpdateSubject(id int, subject model.Subject) error {
	return s.Database.db.UpdateSubject(id, subject)
}

func (s *Services) UpdateTeacher(id int, teacher model.Teacher) error {
	return s.Database.db.UpdateTeacher(id, teacher)
}

func (s *Services) UpdateTeacherSubject(id int, TeacherSubject model.TeacherSubject) error {
	return s.Database.db.UpdateTeacherSubject(id, TeacherSubject)
}

func (s *Services) UpdateRecord(id int, record model.Record) error {
	return s.Database.db.UpdateRecord(id, record)
}

func (s *Services) UpdateTableFile(id int, tableFile model.TableFile) error {
	return s.Database.db.UpdateTableFile(id, tableFile)
}

func (s *Services) UpdateGroupRecords(id int, groupRecords model.GroupRecords) error {
	return s.Database.db.UpdateGroupRecords(id, groupRecords)
}

func (s *Services) UpdatePricingTable(id int, pricingTable model.PricingTable) error {
	return s.Database.db.UpdatePricingTable(id, pricingTable)
}

func (s *Services) UpdatePricingRecord(id int, pricingRecord model.PricingRecord) error {
	return s.Database.db.UpdatePricingRecord(id, pricingRecord)
}

func (s *Services) UpdateTeacherPricingRecord(id int, teacherPricingRecord model.TeacherPricingRecord) error {
	return s.Database.db.UpdateTeacherPricingRecord(id, teacherPricingRecord)
}
