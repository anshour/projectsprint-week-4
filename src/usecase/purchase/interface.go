package purchaseUsecase

import (
	entity "projectsprintw4/src/entities"
	repository "projectsprintw4/src/repositories/purchase"
)

type sPurchaseUsecase struct {
	purchaseRepo repository.PurchaseRepository
}

type PurchaseUsecase interface {
	PurchaseOrder(estimationId string) (orderId string, err error)
	ListNearby(*entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error)
	UserEstimation(*entity.UserEstimationParams, string) (*entity.UserEstimationResult, error)
}

func New(purchaseRepo repository.PurchaseRepository) PurchaseUsecase {
	return &sPurchaseUsecase{purchaseRepo: purchaseRepo}
}
