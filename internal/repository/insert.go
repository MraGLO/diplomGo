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
	_, err = d.db.Exec(context.Background(), "INSERT INTO group(id, specialization_id, name, course, flow_id) VALUES($1, $2, $3, $4, $5)",
		group.ID, group.SpecializationID, group.Name, group.Course, group.FlowID)
	return
}

func (d *DatabaseRepo) InsertSpecialization(specialization *model.Specialization) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO specialization(id, specialization_name, qualification_name) VALUES($1, $2, $3)",
		specialization.ID, specialization.SpecializationName, specialization.QualificationName)
	return
}

func (d *DatabaseRepo) InsertFlow(flow *model.Flow) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO flow(id) VALUES($1)",
		flow.ID)
	return
}

func (d *DatabaseRepo) InsertStudyPlan(studyPlan *model.StudyPlan) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO study_plan(id, flow_id) VALUES($1, $2)",
		studyPlan.ID, studyPlan.FlowID)
	return
}

func (d *DatabaseRepo) InsertRecordStudyPlan(recordStudyPlan *model.RecordStudyPlan) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO record_study_plan(id, semester, subject_id, max_time, independent, consultations, required_time, lectures_and_lessons, practical_exercises, laboratory_exercises, seminars, course_project, intermediate_certification, study_plan_id) VALUES($1, $2)",
		recordStudyPlan.ID, recordStudyPlan.Semester, recordStudyPlan.SubjectID, recordStudyPlan.MaxTime, recordStudyPlan.Independet, recordStudyPlan.Consultations, recordStudyPlan.RequiredTime, recordStudyPlan.LecturesAndLessons, recordStudyPlan.PracticalExercises, recordStudyPlan.LaboratoryExercises, recordStudyPlan.Seminars, recordStudyPlan.CourseProject, recordStudyPlan.IntermediateCertification, recordStudyPlan.StudyPlanID)
	return
}

func (d *DatabaseRepo) InsertLoad(load *model.Load) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO load(id, record_study_plan_id, time_on_subject, teacher_id, group_id) VALUES($1, $2, $3, $4, $5)",
		load.ID, load.RecordStudyPlanID, load.TimeOnSubject, load.TeacherID, load.GroupID)
	return
}
