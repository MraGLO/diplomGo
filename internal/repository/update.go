package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) UpdateSubject(id int, subject model.Subject) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE subject SET name = $1  WHERE id = $2", subject.SubjectName, id)
	return
}

func (d *DatabaseRepo) UpdateTeacher(id int, teacher model.Teacher) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE teacher SET name = $1, surname = $2, patronymic = $3, category = $4  WHERE id = $5",
		teacher.Name, teacher.Surname, teacher.Patronymic, teacher.Category, id)
	return
}

func (d *DatabaseRepo) UpdateTeacherSubject(id int, teacherSubject model.TeacherSubject) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE teacher_subject SET teacher_id = $1, subject_id = $2  WHERE id = $3",
		teacherSubject.TeacherID, teacherSubject.SubjectID, id)
	return
}

func (d *DatabaseRepo) UpdateRecord(id int, record model.Record) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE record SET subject_id = $1, time_semester_one = $2, time_semester_two = $3, teacher_id = $4  WHERE id = $5",
		record.SubjectID, record.TimeSemesterOwn, record.TimeSemesterTwo, record.TeacherID, id)
	return
}

func (d *DatabaseRepo) UpdateTableFile(id int, tableFile model.TableFile) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE table_file SET name = $1  WHERE id = $2",
		tableFile.Name, id)
	return
}

func (d *DatabaseRepo) UpdateGroupRecords(id int, groupRecords model.GroupRecords) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE group_records SET group_id = $1, record_id = $2  WHERE id = $3",
		groupRecords.GroupID, groupRecords.RecordID, id)
	return
}

func (d *DatabaseRepo) UpdatePricingTable(id int, pricingTable model.PricingTable) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE pricing_table SET name = $1  WHERE id = $2",
		pricingTable.Name, id)
	return
}

func (d *DatabaseRepo) UpdatePricingRecord(id int, pricingRecord model.PricingRecord) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE pricing_record SET group_id = $1, first_half_year = $2, second_half_year = $3, theory = $4, practice = $5, LPZ_1 = $6, LPZ_2 = $7, consultation = $8, course_project = $9, hours_first_semester = $10, hours_second_semester = $11, total = $12, pricing_table_id = $13, subject_id = $14  WHERE id = $15",
		pricingRecord.GroupID, pricingRecord.FirstHalfYear, pricingRecord.SecondHalfYear, pricingRecord.Theory, pricingRecord.Practice, pricingRecord.LPZ1, pricingRecord.LPZ2, pricingRecord.Consultation, pricingRecord.CourseProject, pricingRecord.HoursFirstSemester, pricingRecord.HoursSecondSemester, pricingRecord.Total, pricingRecord.PricingTableID, pricingRecord.SubjectID, id)
	return
}

func (d *DatabaseRepo) UpdateTeacherPricingRecord(id int, teacherPricingRecord model.TeacherPricingRecord) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE teacher_pricing_record SET teacher_id = $1, pricing_record_id = $2  WHERE id = $3",
		teacherPricingRecord.TeacherID, teacherPricingRecord.PricingRecordID, id)
	return
}
