package app

import (
	"diplomGo/internal/config"
	"diplomGo/internal/delivery/http"
	"diplomGo/internal/repository"
	"diplomGo/internal/services"
	"diplomGo/pkg/database/postgres"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	pgdb   *pgxpool.Pool
	router *fiber.App
	config *config.Config
}

func (a *App) Init(confPath string) {
	var err error
	a.config, err = config.GetConfig(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// Init router
	a.router = fiber.New()

	a.pgdb, err = postgres.MakePostgres(&a.config.DB)
	if err != nil {
		log.Fatal("Connect db error ", err)
	}

	apirepository := repository.NewRepository(a.pgdb)
	apiservices := services.MakeServices(apirepository)
	handlers := http.MakeHandlers(apiservices)

	InitRoutes(a.router)
	PublicRoutes(a.router, handlers)
	PrivateRoutes(a.router, handlers)
	MiddlewareRoutes(a.router)

}

func (a *App) Run() {
	err := a.router.Listen(a.config.App.Listenport)
	if err != nil {
		log.Println(err)
		return
	}
}

func (a *App) Close() {
	a.pgdb.Close()
}
