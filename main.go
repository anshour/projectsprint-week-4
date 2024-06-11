package main

import (
	"projectsprintw4/src/db"
	v1routes "projectsprintw4/src/http/routes"
	"projectsprintw4/src/utils/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
)

func main() {
	db := db.Init()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Validator = validator.CustomValidator

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := &v1routes.V1Routes{
		Echo: e,
		DB:   db,
	}

	v1.MountAll()

	e.Logger.Fatal(e.Start(":8080"))
}
