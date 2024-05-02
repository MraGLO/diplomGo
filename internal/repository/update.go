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

func (d *DatabaseRepo) UpdateStudyPlan(id int, studyPlan model.StudyPlan) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE study_plan SET flow_id = $1  WHERE id = $2",
		studyPlan.FlowID, id)
	return
}

func (d *DatabaseRepo) UpdateRecordStudyPlan(id int, recordStudyPlan model.RecordStudyPlan) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE record_study_plan SET semester = $1, subject_id = $2, max_time = $3, independent = $4, consultations = $5, required_time = $6, lectures_and_lessons = $7, practical_exercises = $8, laboratory_exercises = $9, seminars = $10, course_project = $11, intermediate_certification = $12, study_plan_id = $13 WHERE id = $14",
		recordStudyPlan.Semester, recordStudyPlan.SubjectID, recordStudyPlan.MaxTime, recordStudyPlan.Independet, recordStudyPlan.Consultations, recordStudyPlan.RequiredTime, recordStudyPlan.LecturesAndLessons, recordStudyPlan.PracticalExercises, recordStudyPlan.LaboratoryExercises, recordStudyPlan.Seminars, recordStudyPlan.CourseProject, recordStudyPlan.IntermediateCertification, recordStudyPlan.StudyPlanID, id)
	return
}

func (d *DatabaseRepo) UpdateLoad(id int, load model.Load) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE load SET record_study_plan_id = $1, time_on_subject = $2, teacher_id = $3, group_id = $4  WHERE id = $5",
		load.RecordStudyPlanID, load.TimeOnSubject, load.TeacherID, load.GroupID, id)
	return
}
