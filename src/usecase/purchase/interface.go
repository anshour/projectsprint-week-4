package purchaseUsecase

import (
	repository "projectsprintw4/src/repositories/purchase"
)

type sPurchaseUsecase struct {
	purchaseRepo repository.PurchaseRepository
}

type PurchaseUsecase interface {
}

func New(purchaseRepo repository.PurchaseRepository) PurchaseUsecase {
	return &sPurchaseUsecase{purchaseRepo: purchaseRepo}
}
