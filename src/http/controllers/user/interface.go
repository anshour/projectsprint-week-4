package userController

import (
	userUsecase "projectsprintw4/src/usecase/user"

	"github.com/labstack/echo/v4"
)

type sUserController struct {
	userUsecase userUsecase.UserUsecase
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Auth interface {
	AdminLogin(c echo.Context) error
	AdminRegister(c echo.Context) error
	UserLogin(c echo.Context) error
	UserRegister(c echo.Context) error
}

func New(usecase userUsecase.UserUsecase) iV1Auth {
	return &sUserController{userUsecase: usecase}
}
