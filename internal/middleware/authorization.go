package middleware

import (
	"fmt"
	"lms-api/internal/abstraction"
	"lms-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type Role string

const (
	IsAdmin    Role = "admin"
	IsStudent  Role = "student"
	IsEmployee Role = "employee"
)

// Authorization checks if context contains at least one given role (OR check) on failure request ends with 401 unauthorized error
func Authorization(allowedRoles ...Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*abstraction.Context)
			roleInContext := cc.Auth.Role
			fmt.Println("==========================================================================================================")
			fmt.Println(roleInContext)
			for _, priv := range allowedRoles {
				if string(priv) == roleInContext {
					return next(c)
				}
			}
			return response.CustomErrorBuilder(405, "You dont have permission", "You dont have permission").Send(c)
		}
	}
}
