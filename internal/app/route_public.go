package app

import (
	"diplomGo/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Static("/subject", "../pages/subject")
	router.Static("/teacher", "../pages/teacher")
	router.Static("/table_file", "../pages/table_file")
	router.Static("/teacher_subject", "../pages/teacher_subject")
	router.Static("/", "../pages/main_page")
	router.Static("/", "../pages")
	router.Static("/pricing", "../pages/pricing")
	router.Static("/pricing_teacher", "../pages/pricing_teacher")
	router.Static("/pricing_teacher_records", "../pages/pricing_teacher_records")
}
