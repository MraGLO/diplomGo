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

	SelectSubjectByID(id int) (subject model.Subject, err error)
	SelectTeacherByID(id int) (teacher model.Teacher, err error)
	SelectTeacherSubjectByID(id int) (teacherSubject model.TeacherSubject, err error)
	SelectGroupByID(id int) (group model.Group, err error)
	SelectSpecializationByID(id int) (specialization model.Specialization, err error)

	SelectIDBySubject(subj string) (id int, err error)

	SelectIDByTeacherSurname(surname string) (id int, err error)

	SelectIDByGroupName(groupName string) (id int, err error)

	InsertSubject(subject *model.Subject) (err error)
	InsertTeacher(teacher *model.Teacher) (err error)
	InsertTeacherSubject(teacherSubject *model.TeacherSubject) (err error)
	InsertGroup(group *model.Group) (err error)
	InsertSpecialization(specialization *model.Specialization) (err error)
	InsertSheet(sheet *model.Sheet) (id int, err error)
	InsertRecord(record *model.Record) (id int, err error)
	InsertSheetRecords(sheetRecord *model.SheetRecords) (err error)
	InsertTableFile(tableFile *model.TableFile) (id int, err error)
	InsertTableFileSheets(tableFIleSheet *model.TableFileSheet) (err error)

	UpdateSubject(id int, subject *model.Subject) (err error)
	UpdateTeacher(id int, teacher model.Teacher) (err error)
	UpdateTeacherSubject(id int, teacherSubject model.TeacherSubject) (err error)

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
