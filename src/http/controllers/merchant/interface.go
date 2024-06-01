package merchantController

import (
	merchantUsecase "projectsprintw4/src/usecase/merchant"

	"github.com/labstack/echo/v4"
)

type sMerchantController struct {
	merchantUsecase merchantUsecase.MerchantUsecase
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponseWithMeta struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Merchant interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	CreateItem(c echo.Context) error
	ListItem(c echo.Context) error
}

func New(merchantUsecase merchantUsecase.MerchantUsecase) iV1Merchant {
	return &sMerchantController{merchantUsecase: merchantUsecase}
}
