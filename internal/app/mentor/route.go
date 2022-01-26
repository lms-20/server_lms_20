package mentor

import (
	"lms-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.POST("", h.Create, middleware.Authentication, middleware.Authorization(middleware.IsAdmin))
	g.PUT("/:id", h.Update, middleware.Authentication, middleware.Authorization(middleware.IsAdmin))
	g.GET("/:id", h.GetByID)
	g.DELETE("/:id", h.Delete, middleware.Authentication, middleware.Authorization(middleware.IsAdmin))
}
