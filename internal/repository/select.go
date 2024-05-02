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
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_id, name, course, flow_id FROM group")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.SpecializationID, &tmp.Name, &tmp.Course, &tmp.FlowID)
		if err != nil {
			return
		}
		groups = append(groups, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectGroupByID(id int) (group model.Group, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, specialization_id, name, course, flow_id  FROM group WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&group.ID, &group.SpecializationID, &group.Name, &group.Course, &group.FlowID)
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

func (d *DatabaseRepo) SelectAllFlows() (flows []model.Flow, err error) {
	var tmp model.Flow
	rows, err := d.db.Query(context.Background(), "SELECT id FROM flow")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID)
		if err != nil {
			return
		}
		flows = append(flows, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectFlowByID(id int) (flow model.Flow, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id FROM flow WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&flow.ID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllStudyPlans() (studyPlans []model.StudyPlan, err error) {
	var tmp model.StudyPlan
	rows, err := d.db.Query(context.Background(), "SELECT id, flow_id FROM study_plan")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.FlowID)
		if err != nil {
			return
		}
		studyPlans = append(studyPlans, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectStudyPlanByID(id int) (studyPlan model.StudyPlan, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, flow_id FROM study_plan WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&studyPlan.ID, &studyPlan.FlowID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllRecordStudyPlans() (recordStudyPlans []model.RecordStudyPlan, err error) {
	var tmp model.RecordStudyPlan
	rows, err := d.db.Query(context.Background(), "SELECT id, semester, subject_id, max_time, independent, consultations, required_time, lectures_and_lessons, practical_exercises, laboratory_exercises, seminars, course_project, intermediate_certification, study_plan_id FROM record_study_plan")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.Semester, &tmp.SubjectID, &tmp.MaxTime, &tmp.Independet, &tmp.Consultations, &tmp.RequiredTime, &tmp.LecturesAndLessons, &tmp.PracticalExercises, &tmp.LaboratoryExercises, &tmp.Seminars, &tmp.CourseProject, &tmp.IntermediateCertification, &tmp.StudyPlanID)
		if err != nil {
			return
		}
		recordStudyPlans = append(recordStudyPlans, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectRecorsStudyPlanByID(id int) (recordStudyPlan model.RecordStudyPlan, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, semester, subject_id, max_time, independent, consultations, required_time, lectures_and_lessons, practical_exercises, laboratory_exercises, seminars, course_project, intermediate_certification, study_plan_id FROM record_study_plan WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&recordStudyPlan.ID, &recordStudyPlan.Semester, &recordStudyPlan.SubjectID, &recordStudyPlan.MaxTime, &recordStudyPlan.Independet, &recordStudyPlan.Consultations, &recordStudyPlan.RequiredTime, &recordStudyPlan.LecturesAndLessons, &recordStudyPlan.PracticalExercises, &recordStudyPlan.LaboratoryExercises, &recordStudyPlan.Seminars, &recordStudyPlan.CourseProject, &recordStudyPlan.IntermediateCertification, &recordStudyPlan.StudyPlanID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllLoads() (loads []model.Load, err error) {
	var tmp model.Load
	rows, err := d.db.Query(context.Background(), "SELECT id, record_study_plan_id, time_on_subject, teacher_id, group_id FROM load")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tmp.ID, &tmp.RecordStudyPlanID, &tmp.TimeOnSubject, &tmp.TeacherID, &tmp.GroupID)
		if err != nil {
			return
		}
		loads = append(loads, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectLoadByID(id int) (load model.Load, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, record_study_plan_id, time_on_subject, teacher_id, group_id FROM load WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&load.ID, &load.RecordStudyPlanID, &load.TimeOnSubject, &load.TeacherID, &load.GroupID)
		if err != nil {
			return
		}
	}
	return
}
