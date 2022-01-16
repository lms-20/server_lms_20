package main

import (
	db "lms-api/database"
	"lms-api/database/migration"
	"lms-api/internal/factory"
	"lms-api/internal/http"
	"lms-api/internal/middleware"
	"lms-api/pkg/elasticsearch"
	"lms-api/pkg/util/env"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv("ENV")
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title lms-api
// @version 0.0.1
// @description This is a doc for lms-api.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv("PORT")

	db.Init()
	migration.Init()
	elasticsearch.Init()

	e := echo.New()
	middleware.Init(e)

	f := factory.NewFactory()
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
