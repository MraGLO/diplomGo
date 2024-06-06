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

	SelectSubjectByID(id int) (subject model.Subject, err error)
	SelectTeacherByID(id int) (teacher model.Teacher, err error)
	SelectTeacherSubjectByID(id int) (teacherSubject model.TeacherSubject, err error)
	SelectGroupByID(id int) (group model.Group, err error)
	SelectRecordByID(id int) (record model.Record, err error)
	SelectTableFilesByID(id int) (tableFile model.TableFile, err error)
	SelectAllTableFiles() (tableFile []model.TableFile, err error)

	SelectAllTableFileGroupsByTableFileID(tableFileID int) (TableFileGroup []model.TableFileGroup, err error)
	SelectAllGroupRecordsByGroupID(groupID int) (groupRecords []model.GroupRecords, err error)

	SelectIDBySubject(subj string) (id int, err error)
	SelectIDByTeacherSurname(surname string) (id int, err error)
	SelectIDByGroupName(groupName string) (id int, err error)

	InsertSubject(subject *model.Subject) (err error)
	InsertTeacher(teacher *model.Teacher) (err error)
	InsertTeacherSubject(teacherSubject *model.TeacherSubject) (err error)
	InsertGroup(group *model.Group) (err error)
	InsertRecord(record *model.Record) (id int, err error)
	InsertGroupRecords(groupRecords *model.GroupRecords) (err error)
	InsertTableFile(tableFile *model.TableFile) (id int, err error)
	InsertTableFileGroups(tableFileGroup *model.TableFileGroup) (err error)

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
