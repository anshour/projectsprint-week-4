package purchaseController

import (
	purchaseUsecase "projectsprintw4/src/usecase/purchase"

	"github.com/labstack/echo/v4"
)

type sPurchaseController struct {
	purchaseUsecase purchaseUsecase.PurchaseUsecase
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Purchase interface {
	PurchaseOrder(c echo.Context) error
	ListOrder(c echo.Context) error
	ListNearby(c echo.Context) error
	UserEstimation(c echo.Context) error
}

func New(purchaseUsecase purchaseUsecase.PurchaseUsecase) iV1Purchase {
	return &sPurchaseController{purchaseUsecase: purchaseUsecase}
}
