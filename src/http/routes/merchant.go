package v1routes

import (
	merchantController "projectsprintw4/src/http/controllers/merchant"
	merchantRepository "projectsprintw4/src/repositories/merchant"
	merchantUsecase "projectsprintw4/src/usecase/merchant"
)

func (i *V1Routes) MountMerchant() {
	repository := merchantRepository.New(i.DB)
	usecase := merchantUsecase.New(repository)
	controller := merchantController.New(usecase)

	//TODO: ADD AUTH MIDLLEWARE
	i.Echo.POST("/admin/merchants", controller.Create)
	i.Echo.GET("/admin/merchants", controller.List)
	i.Echo.POST("/admin/merchants/:merchantId/items", controller.CreateItem)
	i.Echo.GET("/admin/merchants/:merchantId/items", controller.ListItem)

}
