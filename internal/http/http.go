package http

import (
	"fmt"
	docs "lms-api/docs"
	"lms-api/internal/app/category"
	"lms-api/internal/app/courses"
	"lms-api/internal/app/mentor"
	"lms-api/internal/app/sample"
	"lms-api/internal/app/users"
	"lms-api/internal/factory"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
		HOST    = os.Getenv("HOST")
		SCHEME  = os.Getenv("SCHEME")
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	// auth.NewHandler(f).Route(e.Group("/auth"))
	sample.NewHandler(f).Route(e.Group("/samples"))
	users.NewHandler(f).Route(e.Group("/users"))
	mentor.NewHandler(f).Route(e.Group("/mentors"))
	category.NewHandler(f).Route(e.Group("/categories"))
	courses.NewHandler(f).Route(e.Group("/courses"))

}
