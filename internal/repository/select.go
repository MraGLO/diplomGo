package repository

import (
	"context"
	"diplomGo/pkg/model"
)

func (d *DatabaseRepo) SelectAllSubject() (subjects []model.Subject, err error) {
	var tmp model.Subject
	rows, err := d.db.Query(context.Background(), "SELECT id, name FROM subject")
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
	rows, err := d.db.Query(context.Background(), "SELECT id, name FROM subject WHERE id = $1", id)
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

func (d *DatabaseRepo) SelectFirstSubject() (subject model.Subject, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, name FROM subject LIMIT 1")
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
	rows, err := d.db.Query(context.Background(), `SELECT id FROM public."subject" WHERE name = $1`, subj)
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
	rows, err := d.db.Query(context.Background(), "SELECT id, name FROM group")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			return
		}
		groups = append(groups, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectGroupByID(id int) (group model.Group, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id, name FROM public."group" WHERE id = $1`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectFirstGroup() (group model.Group, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id, name FROM public."group" LIMIT 1`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&group.ID, &group.Name)
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

func (d *DatabaseRepo) SelectAllTableFiles() (tableFile []model.TableFile, err error) {
	var tmp model.TableFile
	rows, err := d.db.Query(context.Background(), "SELECT id, name, date FROM table_file")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Name, &tmp.Date)
		if err != nil {
			return
		}
		tableFile = append(tableFile, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectTableFilesByID(id int) (tableFile model.TableFile, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, name, date FROM table_file WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tableFile.ID, &tableFile.Name, &tableFile.Date)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllGroupRecordsByGroupID(groupID int) (groupRecords []model.GroupRecords, err error) {
	var tmp model.GroupRecords
	rows, err := d.db.Query(context.Background(), "SELECT id, group_id, record_id FROM group_records WHERE group_id = $1", groupID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.GroupID, &tmp.RecordID)
		if err != nil {
			return
		}
		groupRecords = append(groupRecords, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectRecordByID(id int) (record model.Record, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, subject_id, time_semester_one, time_semester_two, teacher_id FROM record WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&record.ID, &record.SubjectID, &record.TimeSemesterOwn, &record.TimeSemesterTwo, &record.TeacherID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllTeachersFromTableFileView(tableFileID int) (teachers []model.Teacher, err error) {
	var tmp model.Teacher
	rows, err := d.db.Query(context.Background(), "SELECT distinct teacher_id, teacher_name, teacher_surname, teacher_patronymic, teacher_category FROM public.table_file_view WHERE table_file_id = $1", tableFileID)
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

func (d *DatabaseRepo) SelectAllRecordForPricingFromTableFileView(tableFileID int, teacherID int) (recordForTableFileView []model.RecordForTableFileView, err error) {
	var tmp model.RecordForTableFileView
	rows, err := d.db.Query(context.Background(), "SELECT group_id, group_name, subject_id, subject_name, time_semester_one, time_semester_two FROM public.table_file_view WHERE table_file_id = $1 AND teacher_id = $2", tableFileID, teacherID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.GroupID, &tmp.GroupName, &tmp.SubjectID, &tmp.SubjectName, &tmp.TimeSemesterOwn, &tmp.TimeSemesterTwo)
		if err != nil {
			return
		}
		recordForTableFileView = append(recordForTableFileView, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectSubjectsFromTeacherSubjectViewByTeacherID(teacherID int) (subjects []model.Subject, err error) {
	var tmp model.Subject
	rows, err := d.db.Query(context.Background(), "SELECT subject_id, subject_name FROM public.teacher_subject_view WHERE teacher_id = $1", teacherID)
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

func (d *DatabaseRepo) SelectAllPricingTable() (pricing []model.PricingTable, err error) {
	var tmp model.PricingTable
	rows, err := d.db.Query(context.Background(), "SELECT id, name, date FROM pricing_table")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Name, &tmp.Date)
		if err != nil {
			return
		}
		pricing = append(pricing, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectAllTeachersFromPricingTableView(pricinTableID int) (teachers []model.Teacher, err error) {
	var tmp model.Teacher
	rows, err := d.db.Query(context.Background(), "SELECT distinct teacher_id, teacher_name, teacher_surname, teacher_patronymic, teacher_category FROM public.pricing_table_view WHERE pricing_table_id = $1", pricinTableID)
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

func (d *DatabaseRepo) SelectAllTeacherPricingRecordFromPricingTAbleViewByData(pricingTeacherStruct model.GetPricingTeacherStruct) (teacherPricingRecord []model.TeacherPricingRecord, err error) {
	var tmp model.TeacherPricingRecord
	rows, err := d.db.Query(context.Background(), "SELECT teacher_id, id FROM pricing_table_view WHERE teacher_id = $1 AND pricing_table_id = $2", pricingTeacherStruct.TeacherID, pricingTeacherStruct.PricingID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.TeacherID, &tmp.PricingRecordID)
		if err != nil {
			return
		}
		teacherPricingRecord = append(teacherPricingRecord, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectAllPricingRecordFromPricingTAbleViewByData(pricingTeacherStruct model.GetPricingTeacherStruct) (pricingRecordView []model.PricingRecordView, err error) {
	var tmp model.PricingRecordView
	rows, err := d.db.Query(context.Background(), `SELECT id, group_name, first_half_year, second_half_year, theory, practice, "LPZ_1", "LPZ_2", consultation, course_project, hours_first_semester, hours_second_semester, total, subject_name FROM pricing_table_view WHERE teacher_id = $1 AND pricing_table_id = $2`, pricingTeacherStruct.TeacherID, pricingTeacherStruct.PricingID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Group, &tmp.FirstHalfYear, &tmp.SecondHalfYear, &tmp.Theory, &tmp.Practice, &tmp.LPZ1, &tmp.LPZ2, &tmp.Consultation, &tmp.CourseProject, &tmp.HoursFirstSemester, &tmp.HoursSecondSemester, &tmp.Total, &tmp.Subject)
		if err != nil {
			return
		}
		pricingRecordView = append(pricingRecordView, tmp)
	}
	return
}
