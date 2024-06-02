package middleware

import (
	"net/http"
	"projectsprintw4/src/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}

			tokenParts := strings.Split(tokenString, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Invalid Authorization Header",
				})

			}

			token := tokenParts[1]

			payload, err := jwt.Verify(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Status:  false,
					Message: "Invalid Authorization Header",
				})
			}
			println("user: ", payload.UserId)
			c.Set("userId", payload.UserId)
			c.Set("role", payload.Role)

			return next(c)
		}
	}
}
