package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) InsertSubject(subject *model.Subject) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO subject(name) VALUES($1)", subject.SubjectName)
	return
}

func (d *DatabaseRepo) InsertTeacher(teacher *model.Teacher) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO teacher(name, surname, patronymic, category) VALUES($1, $2, $3, $4)",
		teacher.Name, teacher.Surname, teacher.Patronymic, teacher.Category)
	return
}

func (d *DatabaseRepo) InsertTeacherSubject(teacherSubject *model.TeacherSubject) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO teacher_subject(teacher_id, subject_id) VALUES($1, $2)",
		teacherSubject.TeacherID, teacherSubject.SubjectID)
	return
}

func (d *DatabaseRepo) InsertGroup(group *model.Group) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO group(id, name) VALUES($1, $2)",
		group.ID, group.Name)
	return
}

func (d *DatabaseRepo) InsertRecord(record *model.Record) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO record(subject_id, time_semester_one, time_semester_two, teacher_id, table_file_id) VALUES($1, $2, $3, $4, $5) RETURNING id",
		record.SubjectID, record.TimeSemesterOwn, record.TimeSemesterTwo, record.TeacherID, record.TableFileID)
	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertTableFile(tableFile *model.TableFile) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO table_file(name, date) VALUES($1, $2) RETURNING id",
		tableFile.Name, tableFile.Date)

	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertGroupRecords(groupRecords *model.GroupRecords) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO group_records(group_id, record_id) VALUES($1, $2)",
		groupRecords.GroupID, groupRecords.RecordID)
	return
}

func (d *DatabaseRepo) InsertPricingTable(pricingTable *model.PricingTable) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO pricing_table(name, date) VALUES($1, $2) RETURNING id",
		pricingTable.Name, pricingTable.Date)

	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertPricingRecord(pricingRecord *model.PricingRecord) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO pricing_record(group_id, first_half_year, second_half_year, theory, practice, LPZ_1, LPZ_2, consultation, course_project, hours_first_semester, hours_second_semester, total, pricing_table_id, subject_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
		pricingRecord.GroupID, pricingRecord.FirstHalfYear, pricingRecord.SecondHalfYear, pricingRecord.Theory, pricingRecord.Practice, pricingRecord.LPZ1, pricingRecord.LPZ2, pricingRecord.Consultation, pricingRecord.CourseProject, pricingRecord.HoursFirstSemester, pricingRecord.HoursSecondSemester, pricingRecord.Total, pricingRecord.PricingTableID, pricingRecord.SubjectID)

	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertTeacherPricingRecord(teacherPricingRecord *model.TeacherPricingRecord) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO teacher_pricing_record(teacher_id, pricing_record_id) VALUES($1, $2)",
		teacherPricingRecord.TeacherID, teacherPricingRecord.PricingRecordID)
	return
}
