package v1routes

import (
	userController "projectsprintw4/src/http/controllers/user"
	userRepository "projectsprintw4/src/repositories/user"
	userUsecase "projectsprintw4/src/usecase/user"
)

func (i *V1Routes) MountUser() {
	userRepository := userRepository.New(i.DB)
	userUseCase := userUsecase.New(userRepository)
	controller := userController.New(userUseCase)

	i.Echo.POST("/admin/login", controller.AdminLogin)
	i.Echo.POST("/admin/register", controller.AdminRegister)
	i.Echo.POST("/users/login", controller.UserLogin)
	i.Echo.POST("/users/register", controller.UserRegister)
}
