package purchaseUsecase

import entity "projectsprintw4/src/entities"

func (uc *sPurchaseUsecase) ListNearby(p *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {
	items, err := uc.purchaseRepo.ListAllNearby(p)

	if err != nil {
		empty := make([]entity.ListNearbyMerchantResult, 0)
		return &empty, err
	}

	return items, nil
}
