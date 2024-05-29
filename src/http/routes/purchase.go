package v1routes

import (
	"net/http"
	purchaseController "projectsprintw4/src/http/controllers/purchase"
	middleware "projectsprintw4/src/http/middlewares"
	purchaseRepository "projectsprintw4/src/repositories/purchase"
	purchaseUsecase "projectsprintw4/src/usecase/purchase"

	"github.com/labstack/echo/v4"
)

func (i *V1Routes) MountPurchase() {
	repository := purchaseRepository.New(i.DB)
	usecase := purchaseUsecase.New(repository)
	controller := purchaseController.New(usecase)
	println(controller)
	//TODO: ADD AUTH MIDLLEWARE
	i.Echo.GET("/merchants/nearby/:latlong", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, middleware.Authentication())

	g := i.Echo.Group("/users")
	g.Use(middleware.Authentication())

	g.GET("/estimate", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	g.POST("/orders", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	g.GET("/orders", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}
