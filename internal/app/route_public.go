package app

import (
	"diplomGo/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Static("/", "../SubjectPage")
	router.Static("/", "../FilePage")
	router.Static("/", "../PricingPage")
	router.Static("/", "../MainPage")
}
