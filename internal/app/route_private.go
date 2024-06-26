package app

import (
	"diplomGo/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Get("subject/all", handlers.GetAllSubjects)
	router.Get("subject/:subjectID", handlers.GetSubjectById)
	router.Get("teacher/all", handlers.GetAllTeachers)
	router.Get("teacher/:teacherID", handlers.GetTeacherById)
	router.Get("teacherSubject/all", handlers.GetAllTeacherSubject)
	router.Get("teacherSubject/:teacherSubjectID", handlers.GetTeacherSubjectById)
	router.Get("createPricing/:tableFileID", handlers.CreatePricing)
	router.Get("tableFile/all", handlers.GetAllTableFile)
	router.Get("teacher_subject/:teacherID", handlers.GetTeacherSubjectViewById)
	router.Get("pricing/all", handlers.GetAllPricing)
	router.Get("pricingTeacher/:pricingID", handlers.GetTeacherByPricingId)

	subject := router.Group("/subject")
	teacher := router.Group("/teacher")
	teacherSubject := router.Group("/teacherSubject")
	tableFile := router.Group("/tableFile")
	pricing := router.Group("/pricing")
	pricingTeacher := router.Group("/pricingTeacher")
	pricingTeacherRecord := router.Group("/pricingTeacherRecord")

	tableFile.Post("/add", handlers.AddFile)
	tableFile.Delete("delete/:TableFileID", handlers.DeleteTableFile)

	subject.Post("/add", handlers.AddSubject)
	subject.Put("/update/:subjectID", handlers.UpdateSubject)
	subject.Delete("/delete/:subjectID", handlers.DeleteSubject)

	teacher.Post("/add", handlers.AddTeacher)
	teacher.Put("/update/:teacherID", handlers.UpdateTeacher)
	teacher.Delete("/delete/:teacherID", handlers.DeleteTeacher)

	teacherSubject.Post("/add", handlers.AddTeacherSubjects)
	teacherSubject.Put("/update/:teacherSubjectID", handlers.UpdateTeacherSubject)
	teacherSubject.Delete("/delete/:teacherSubjectID", handlers.DeleteTeacherSubject)
	teacherSubject.Delete("/delete", handlers.DeleteTeacherSubjectByData)

	pricing.Post("/add", handlers.CreatePricing)
	pricing.Put("/update/:pricingTableID", handlers.UpdatePricing)
	pricing.Delete("/delete/:pricingTableID", handlers.DeletePricing)

	pricingTeacher.Post("/add", handlers.AddPricingTeacher)
	pricingTeacher.Delete("/delete", handlers.DeleteTeacherPricingByTeacherId)

	pricingTeacherRecord.Post("/all", handlers.GetTeacherPricingRecords)
}
