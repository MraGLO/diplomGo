package repository

import "context"

func (d *DatabaseRepo) DeleteSubject(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM subject WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

//Созадать каскадное удаление с таблицей teacher_subject
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
