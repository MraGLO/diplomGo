package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) DeleteSubject(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM subject WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTeacher(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM teacher WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTeacherSubject(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM teacher_subject WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteRecord(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM record WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTableFile(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM table_file WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteGroupRecords(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM group_records WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeletePricingTable(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM pricing_table WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeletePricingRecord(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM pricing_record WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTeacherPricingRecord(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM teacher_pricing_record WHERE id = $1", id)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTeacherPricingRecordByData(teacherSubject model.TeacherSubject) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM teacher_subject WHERE teacher_id = $1 AND subject_id = $2", teacherSubject.TeacherID, teacherSubject.SubjectID)
	if err != nil {
		return
	}
	count = int(com.RowsAffected())
	return
}
