package repository

import (
	"diplomGo/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	SelectAllSubject() (subjects []model.Subject, err error)
	SelectAllTeachers() (teachers []model.Teacher, err error)
	SelectAllTeacherSubject() (teacherSubject []model.TeacherSubject, err error)
	SelectAllGroups() (groups []model.Group, err error)
	SelectAllSpecializations() (specializations []model.Specialization, err error)
	SelectAllFlows() (flows []model.Flow, err error)
	SelectAllStudyPlans() (studyPlans []model.StudyPlan, err error)
	SelectAllRecordStudyPlans() (recordStudyPlans []model.RecordStudyPlan, err error)
	SelectAllLoads() (loads []model.Load, err error)

	SelectSubjectByID(id int) (subject model.Subject, err error)
	SelectTeacherByID(id int) (teacher model.Teacher, err error)
	SelectTeacherSubjectByID(id int) (teacherSubject model.TeacherSubject, err error)
	SelectGroupByID(id int) (group model.Group, err error)
	SelectSpecializationByID(id int) (specialization model.Specialization, err error)
	SelectFlowByID(id int) (flow model.Flow, err error)
	SelectStudyPlanByID(id int) (studyPlan model.StudyPlan, err error)
	SelectRecordStudyPlanByID(id int) (recordStudyPlan model.RecordStudyPlan, err error)
	SelectLoadByID(id int) (load model.Load, err error)

	InsertSubject(subject *model.Subject) (err error)
	InsertTeacher(teacher *model.Teacher) (err error)
	InsertTeacherSubject(teacherSubject *model.TeacherSubject) (err error)
	InsertGroup(group *model.Group) (err error)
	InsertSpecialization(specialization *model.Specialization) (err error)
	InsertFlow(flow *model.Flow) (err error)
	InsertStudyPlan(studyPlan *model.StudyPlan) (err error)
	InsertRecordStudyPlan(recordStudyPlan *model.RecordStudyPlan) (err error)
	InsertLoad(load *model.Load) (err error)

	UpdateSubject(id int, subject *model.Subject) (err error)
	UpdateTeacher(id int, teacher model.Teacher) (err error)
	UpdateTeacherSubject(id int, teacherSubject model.TeacherSubject) (err error)
	UpdateStudyPlan(id int, studyPlan model.StudyPlan) (err error)
	UpdateRecordStudyPlan(id int, recordStudyPlan model.RecordStudyPlan) (err error)
	UpdateLoad(id int, load model.Load) (err error)

	DeleteSubject(id int) (count int, err error)
	DeleteTeacher(id int) (count int, err error)
	DeleteTeacherSubject(id int) (count int, err error)
}

type Repository struct {
	Database Database
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Database: newDatabaseRepo(db)}
}
