package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kimcodec/test_api_task/controllers"
	"github.com/kimcodec/test_api_task/internal/repository"
	"github.com/kimcodec/test_api_task/internal/services"
	openapi "github.com/kimcodec/test_api_task/outer_api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
)

const (
	defaultAddress = ":8080"
	defaultDBURI   = "postgres://postgres:postgres@localhost:5432/test_api?sslmode=disable"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := sqlx.Connect("postgres", defaultDBURI)
	if err != nil {
		log.Fatal("[FATAL] Failed to connect to db: ", err.Error())
	}

	oc := repository.NewOwnerRepository(db)
	cr := repository.NewCarRepository(db)

	outerApi := openapi.NewAPIClient(openapi.NewConfiguration())

	cs := services.NewCarService(cr, oc, outerApi)

	e := echo.New()
	defer e.Close()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	controllers.NewCarController(e, cs)

	if err := e.Start(defaultAddress); err != nil {
		log.Fatal("[FATAL] Failed to start app: ", err.Error())
	}
}
