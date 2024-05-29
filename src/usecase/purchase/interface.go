package purchaseUsecase

import (
	entity "projectsprintw4/src/entities"
	repository "projectsprintw4/src/repositories/purchase"
)

type sPurchaseUsecase struct {
	purchaseRepo repository.PurchaseRepository
}

type PurchaseUsecase interface {
	ListNearby(*entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error)
}

func New(purchaseRepo repository.PurchaseRepository) PurchaseUsecase {
	return &sPurchaseUsecase{purchaseRepo: purchaseRepo}
}
