package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) UpdateSubject(id int, subject *model.Subject) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE subject SET subject_name = $1  WHERE id = $2", subject.SubjectName, id)
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
