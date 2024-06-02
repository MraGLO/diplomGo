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
	router.Get("specialization/all", handlers.GetAllSpecializations)
	router.Get("specialization/:specializationID", handlers.GetSpecializationById)

	subject := router.Group("/subject")
	teacher := router.Group("/teacher")
	teacherSubject := router.Group("/teacherSubject")
	file := router.Group("/file")

	file.Post("/add", handlers.AddFile)

	subject.Post("/add", handlers.AddSubject)
	subject.Put("/:subjectID", handlers.UpdateSubject)
	subject.Delete("/:subjectID", handlers.DeleteSubject)

	teacher.Post("/add", handlers.AddTeacher)
	teacher.Put("/:teacherID", handlers.UpdateTeacher)
	teacher.Delete("/:teacherID", handlers.DeleteTeacher)

	teacherSubject.Post("/add", handlers.AddTeacherSubjects)
	teacherSubject.Put("/:teacherSubjectID", handlers.UpdateTeacherSubject)
	teacherSubject.Delete("/:teacherSubjectID", handlers.DeleteTeacherSubject)

}
