package v1routes

import (
	"net/http"
	purchaseController "projectsprintw4/src/http/controllers/purchase"
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
	i.Echo.POST("/merchants/nearby/:lat,long", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	i.Echo.GET("/users/estimate", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	i.Echo.POST("/users/orders", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	i.Echo.GET("/users/orders", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}
