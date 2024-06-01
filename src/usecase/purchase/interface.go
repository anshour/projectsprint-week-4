package purchaseUsecase

import (
	entity "projectsprintw4/src/entities"
	repository "projectsprintw4/src/repositories/purchase"
)

type sPurchaseUsecase struct {
	purchaseRepo repository.PurchaseRepository
}

type PurchaseUsecase interface {
	ListNearby(*entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error)
	UserEstimation(*entity.UserEstimationParams) (*entity.UserEstimationResult, error)
}

func New(purchaseRepo repository.PurchaseRepository) PurchaseUsecase {
	return &sPurchaseUsecase{purchaseRepo: purchaseRepo}
}
