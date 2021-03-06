package middleware

import (
	"fmt"
	"os"
	"strings"

	"lms-api/internal/abstraction"
	res "lms-api/pkg/util/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// type Role struct {
// 	Roles []string
// }

// type Privilege string

// const (
// 	IsAdmin    Privilege = "admin"
// 	IsStudent  Privilege = "student"
// 	IsEmployee Privilege = "employee"
// )

// // AuthContext holds information about user who has been authenticated for request
// type AuthContext struct {
// 	Privileges []Privilege
// }

// func Authorization(privileges ...Privilege) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			// raw := c.Get(SecurityContextKey)
// 			// fmt.Println("==============================================================")
// 			// fmt.Println(raw)
// 			// if raw == nil {
// 			// 	return echo.ErrUnauthorized
// 			// }
// 			cc := c.(*abstraction.Context)
// 			currentRole := cc.Auth.Role
// 			success := checkRole(currentRole)
// 			if cc != nil {
// 				return echo.ErrUnauthorized
// 			}
// 			return next(c)
// 		}
// 	}
// }

// func checkRole(roless []string, currentRole string) bool {
// 	role := strings.Join(roless[:], " ")
// 	return strings.Contains(role, currentRole)
// }

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}

		var id int
		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID != nil {
			id = int(destructID.(float64))
		} else {
			id = 0
		}

		var name string
		destructName := token.Claims.(jwt.MapClaims)["name"]
		if destructName != nil {
			name = destructName.(string)
		} else {
			name = ""
		}

		var email string
		destructEmail := token.Claims.(jwt.MapClaims)["email"]
		if destructEmail != nil {
			email = destructEmail.(string)
		} else {
			email = ""
		}

		var role string
		destructRole := token.Claims.(jwt.MapClaims)["role"]
		if destructRole != nil {
			role = destructRole.(string)
		} else {
			role = ""
		}

		cc := c.(*abstraction.Context)
		cc.Auth = &abstraction.AuthContext{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  role,
		}

		return next(cc)
	}
}
