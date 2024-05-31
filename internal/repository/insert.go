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
