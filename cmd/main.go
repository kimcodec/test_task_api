package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/kimcodec/test_api_task/controllers"
	_ "github.com/kimcodec/test_api_task/docs"
	openapi "github.com/kimcodec/test_api_task/internal/outer_api"
	"github.com/kimcodec/test_api_task/internal/repository"
	"github.com/kimcodec/test_api_task/internal/services"

	"fmt"
	"os"
)

const (
	defaultPort  = "8080"
	defaultDBURI = "postgres://postgres:postgres@postgres:5432/test_api_db?sslmode=disable"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title		API
//	@version	1

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	dbURI := ""
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || sslMode == "" {
		log.Println("[ERROR] Failed to make db URI. Using default value")
		dbURI = defaultDBURI
	} else {
		dbURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	}

	db, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Fatal("[FATAL] Failed to connect to db: ", err.Error())
	}

	oc := repository.NewOwnerRepository(db)
	cr := repository.NewCarRepository(db)

	outerApi := openapi.NewAPIClient(openapi.NewConfiguration())

	cs := services.NewCarService(cr, oc, outerApi)

	e := echo.New()
	defer e.Close()

	logger := log.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP: true,
		LogHost:     true,
		LogProtocol: true,
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logger.WithFields(log.Fields{
				"RemoteIP": values.RemoteIP,
				"Host":     values.Host,
				"Protocol": values.Protocol,
				"Method":   values.Method,
				"URI":      values.URI,
				"status":   values.Status,
			}).Info("request details")

			return nil
		},
	}))
	e.Use(middleware.CORS())

	controllers.NewCarController(e, cs)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Println("[ERROR] Failed to get app_port. Using default value.")
		appPort = defaultPort
	}
	if err := e.Start(":" + appPort); err != nil {
		log.Fatal("[FATAL] Failed to start app: ", err.Error())
	}
}
