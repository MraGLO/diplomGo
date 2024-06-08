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

// func (d *DatabaseRepo) InsertTableFileGroups(tableFileGroup *model.TableFileGroup) (err error) {
// 	_, err = d.db.Exec(context.Background(), "INSERT INTO table_file_groups(group_id, table_file_id) VALUES($1, $2)",
// 		tableFileGroup.GroupID, tableFileGroup.TableFileID)
// 	return
// }

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
