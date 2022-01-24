package http

import (
	"fmt"
	docs "lms-api/docs"
	"lms-api/internal/app/category"
	"lms-api/internal/app/chapter"
	"lms-api/internal/app/courses"
	"lms-api/internal/app/lessons"
	"lms-api/internal/app/mentor"
	"lms-api/internal/app/mycourse"
	"lms-api/internal/app/note"
	"lms-api/internal/app/order"
	"lms-api/internal/app/review"
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
	sample.NewHandler(f).Route(e.Group("/samples"))
	users.NewHandler(f).Route(e.Group("/users"))
	mentor.NewHandler(f).Route(e.Group("/mentors"))
	category.NewHandler(f).Route(e.Group("/categories"))
	courses.NewHandler(f).Route(e.Group("/courses"))
	chapter.NewHandler(f).Route(e.Group("/chapters"))
	lessons.NewHandler(f).Route(e.Group("/lessons"))
	note.NewHandler(f).Route(e.Group("/notes"))
	review.NewHandler(f).Route(e.Group("/reviews"))
	mycourse.NewHandler(f).Route(e.Group("/mycourses"))
	order.NewHandler(f).Route(e.Group("/orders"))

}
