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
	router.Get("studyPlan/all", handlers.GetAllStudyPlans)
	router.Get("StudyPlan/:studyPlanID", handlers.GetStudyPlanById)
	router.Get("recordStudyPlan/all", handlers.GetAllRecordStudyPlans)
	router.Get("recordStudyPlan/:recordStudyPlanID", handlers.GetRecordStudyPlanById)
	router.Get("load/all", handlers.GetAllLoad)
	router.Get("load/:loadID", handlers.GetLoadById)

	subject := router.Group("/subject")
	teacher := router.Group("/teacher")
	teacherSubject := router.Group("/teacherSubject")
	//specialization :=router.Group("/specialization")
	studyPlan := router.Group("/studyPlan")
	recordStudyPlan := router.Group("/recordStudyPlan")
	load := router.Group("/load")

	studyPlan.Post("/add", handlers.AddStudyPlan)

	recordStudyPlan.Post("/add", handlers.AddRecordStudyPlan)

	load.Post("/add", handlers.AddLoad)

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
