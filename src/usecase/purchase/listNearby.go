package purchaseUsecase

import entity "projectsprintw4/src/entities"

func (uc *sPurchaseUsecase) ListNearby(p *entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error) {
	items, err := uc.purchaseRepo.ListAllNearby(p)

	if err != nil {
		empty := make([]entity.ListNearbymerchantFinalResult, 0)
		return &empty, err
	}

	return items, nil
}
