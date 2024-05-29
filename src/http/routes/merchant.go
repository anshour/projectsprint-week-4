package v1routes

import (
	merchantController "projectsprintw4/src/http/controllers/merchant"
	middleware "projectsprintw4/src/http/middlewares"
	merchantRepository "projectsprintw4/src/repositories/merchant"
	merchantUsecase "projectsprintw4/src/usecase/merchant"
)

func (i *V1Routes) MountMerchant() {
	repository := merchantRepository.New(i.DB)
	usecase := merchantUsecase.New(repository)
	controller := merchantController.New(usecase)

	//TODO: ADD AUTH MIDLLEWARE
	g := i.Echo.Group("/admin")
	g.Use(middleware.Authentication())
	g.POST("/merchants", controller.Create)
	g.GET("/merchants", controller.List)
	g.POST("/merchants/:merchantId/items", controller.CreateItem)
	g.GET("/merchants/:merchantId/items", controller.ListItem)

}
