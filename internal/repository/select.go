package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) SelectAllSubject() (subjects []model.Subject, err error) {
	var tmp model.Subject
	rows, err := d.db.Query(context.Background(), "SELECT id, subject_name FROM subject")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.SubjectName)
		if err != nil {
			return
		}
		subjects = append(subjects, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectSubjectByID(id int) (subject model.Subject, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, subject_name FROM subject WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&subject.ID, &subject.SubjectName)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectIDBySubject(subj string) (id int, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id FROM public."subject" WHERE subject_name = $1`, subj)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllTeachers() (teachers []model.Teacher, err error) {
	var tmp model.Teacher
	rows, err := d.db.Query(context.Background(), "SELECT id, name, surname, patronymic, category FROM teacher")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Name, &tmp.Surname, &tmp.Patronymic, &tmp.Category)
		if err != nil {
			return
		}
		teachers = append(teachers, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectTeacherByID(id int) (teacher model.Teacher, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, name, surname, patronymic, category FROM teacher WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Patronymic, &teacher.Category)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectIDByTeacherSurname(surname string) (id int, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id FROM public."teacher" WHERE surname = $1`, surname)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllTeacherSubject() (teacherSubject []model.TeacherSubject, err error) {
	var tmp model.TeacherSubject
	rows, err := d.db.Query(context.Background(), "SELECT id, teacher_id, subject_id FROM teacher_subject")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.TeacherID, &tmp.SubjectID)
		if err != nil {
			return
		}
		teacherSubject = append(teacherSubject, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectTeacherSubjectByID(id int) (teacherSubject model.TeacherSubject, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, teacher_id, subject_id FROM teacher_subject WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&teacherSubject.ID, &teacherSubject.TeacherID, &teacherSubject.SubjectID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllGroups() (groups []model.Group, err error) {
	var tmp model.Group
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_id, name, course FROM group")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.SpecializationID, &tmp.Name, &tmp.Course)
		if err != nil {
			return
		}
		groups = append(groups, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectGroupByID(id int) (group model.Group, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_id, name, course  FROM group WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&group.ID, &group.SpecializationID, &group.Name, &group.Course)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllSpecializations() (specializations []model.Specialization, err error) {
	var tmp model.Specialization
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_name, qualification_name FROM specialization")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.SpecializationName, &tmp.QualificationName)
		if err != nil {
			return
		}
		specializations = append(specializations, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectSpecializationByID(id int) (specialization model.Specialization, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_name, qualification_name FROM specialization WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&specialization.ID, &specialization.SpecializationName, &specialization.QualificationName)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectIDByGroupName(groupName string) (id int, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id FROM public."group" WHERE name = $1`, groupName)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
	}
	return
}
