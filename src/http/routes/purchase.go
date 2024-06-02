package v1routes

import (
	purchaseController "projectsprintw4/src/http/controllers/purchase"
	middleware "projectsprintw4/src/http/middlewares"
	purchaseRepository "projectsprintw4/src/repositories/purchase"
	purchaseUsecase "projectsprintw4/src/usecase/purchase"
)

func (i *V1Routes) MountPurchase() {
	repository := purchaseRepository.New(i.DB)
	usecase := purchaseUsecase.New(repository)
	controller := purchaseController.New(usecase)

	i.Echo.GET("/merchants/nearby/:latlong", controller.ListNearby)

	g := i.Echo.Group("/users")

	g.Use(middleware.Authentication())
	g.POST("/estimate", controller.UserEstimation)

	g.POST("/orders", controller.PurchaseOrder)
	g.GET("/orders", controller.ListOrder)

}
