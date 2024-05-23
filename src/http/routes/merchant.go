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

	i.Echo.POST("/admin/merchants", controller.Create)
}
