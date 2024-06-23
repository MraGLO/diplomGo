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
	SelectAllPricingTable() (pricing []model.PricingTable, err error)

	SelectSubjectByID(id int) (subject model.Subject, err error)
	SelectTeacherByID(id int) (teacher model.Teacher, err error)
	SelectTeacherSubjectByID(id int) (teacherSubject model.TeacherSubject, err error)
	SelectGroupByID(id int) (group model.Group, err error)
	SelectRecordByID(id int) (record model.Record, err error)
	SelectTableFilesByID(id int) (tableFile model.TableFile, err error)
	SelectAllTableFiles() (tableFile []model.TableFile, err error)

	SelectAllGroupRecordsByGroupID(groupID int) (groupRecords []model.GroupRecords, err error)
	SelectSubjectsFromTeacherSubjectViewByTeacherID(teacherID int) (subjects []model.Subject, err error)
	SelectAllTeacherPricingRecordFromPricingTAbleViewByData(pricingTeacherStruct model.GetPricingTeacherStruct) (teacherPricingRecord []model.TeacherPricingRecord, err error)

	SelectIDBySubject(subj string) (id int, err error)
	SelectIDByTeacherSurname(surname string) (id int, err error)
	SelectIDByGroupName(groupName string) (id int, err error)

	SelectAllTeachersFromTableFileView(tableFileID int) (teachers []model.Teacher, err error)
	SelectAllRecordForPricingFromTableFileView(tableFileID int, teacherID int) (recordForTableFileView []model.RecordForTableFileView, err error)
	SelectAllTeachersFromPricingTableView(pricinTableID int) (teachers []model.Teacher, err error)

	SelectFirstSubject() (subject model.Subject, err error)
	SelectFirstGroup() (group model.Group, err error)

	InsertSubject(subject *model.Subject) (err error)
	InsertTeacher(teacher *model.Teacher) (err error)
	InsertTeacherSubject(teacherSubject *model.TeacherSubject) (err error)
	InsertGroup(group *model.Group) (err error)
	InsertRecord(record *model.Record) (id int, err error)
	InsertGroupRecords(groupRecords *model.GroupRecords) (err error)
	InsertTableFile(tableFile *model.TableFile) (id int, err error)
	InsertPricingTable(pricingTable *model.PricingTable) (id int, err error)
	InsertTeacherPricingRecord(teacherPricingRecord *model.TeacherPricingRecord) (err error)
	InsertPricingRecord(pricingRecord *model.PricingRecord) (id int, err error)

	UpdateSubject(id int, subject model.Subject) (err error)
	UpdateTeacher(id int, teacher model.Teacher) (err error)
	UpdateTeacherSubject(id int, teacherSubject model.TeacherSubject) (err error)
	UpdateRecord(id int, record model.Record) (err error)
	UpdateTableFile(id int, tableFile model.TableFile) (err error)
	UpdateGroupRecords(id int, groupRecords model.GroupRecords) (err error)
	UpdatePricingTable(id int, pricingTable model.PricingTable) (err error)
	UpdatePricingRecord(id int, pricingRecord model.PricingRecord) (err error)
	UpdateTeacherPricingRecord(id int, teacherPricingRecord model.TeacherPricingRecord) (err error)

	DeleteSubject(id int) (count int, err error)
	DeleteTeacher(id int) (count int, err error)
	DeleteTeacherSubject(id int) (count int, err error)
	DeleteRecord(id int) (count int, err error)
	DeleteTableFile(id int) (count int, err error)
	DeleteGroupRecords(id int) (count int, err error)
	DeletePricingTable(id int) (count int, err error)
	DeletePricingRecord(id int) (count int, err error)
	DeleteTeacherPricingRecord(id int) (count int, err error)
	DeleteTeacherSubjectByData(teacherSubject model.TeacherSubject) (count int, err error)
	DeleteTeacherPricingRecordByTeacherID(id int) (count int, err error)
}

type Repository struct {
	Database Database
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Database: newDatabaseRepo(db)}
}
