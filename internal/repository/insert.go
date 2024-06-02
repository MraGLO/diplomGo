package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) InsertSubject(subject *model.Subject) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO subject(subject_name) VALUES($1)", subject.SubjectName)
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
	_, err = d.db.Exec(context.Background(), "INSERT INTO group(id, specialization_id, name, course) VALUES($1, $2, $3, $4)",
		group.ID, group.SpecializationID, group.Name, group.Course)
	return
}

func (d *DatabaseRepo) InsertSpecialization(specialization *model.Specialization) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO specialization(id, specialization_name, qualification_name) VALUES($1, $2, $3)",
		specialization.ID, specialization.SpecializationName, specialization.QualificationName)
	return
}

func (d *DatabaseRepo) InsertSheet(sheet *model.Sheet) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO sheet(group_id) VALUES($1)",
		sheet.GroupID)
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertRecord(record *model.Record) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO record(subject_id, time_semester_own, time_semester_two, teacher_id) VALUES($1, $2, $3, $4)",
		record.SubjectID, record.TimeSemesterOwn, record.TimeSemesterTwo, record.TeacherID)
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertTableFile(tableFile *model.TableFile) (id int, err error) {
	row, err := d.db.Query(context.Background(), "INSERT INTO table_file(name, date) VALUES($1, $2)",
		tableFile.Name, tableFile.Date)
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) InsertSheetRecords(sheetRecords *model.SheetRecords) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO sheet_records(sheet_id, record_id) VALUES($1, $2)",
		sheetRecords.SheetID, sheetRecords.RecordID)
	return
}

func (d *DatabaseRepo) InsertTableFileSheets(tableFileSheet *model.TableFileSheet) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO table_file_sheets(sheet_id, table_file_id) VALUES($1, $2)",
		tableFileSheet.SheetID, tableFileSheet.TableFileID)
	return
}
