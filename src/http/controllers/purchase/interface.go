package purchaseController

import (
	purchaseUsecase "projectsprintw4/src/usecase/purchase"
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
}

func New(purchaseUsecase purchaseUsecase.PurchaseUsecase) iV1Purchase {
	return &sPurchaseController{purchaseUsecase: purchaseUsecase}
}
