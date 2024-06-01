package purchaseUsecase

import entity "projectsprintw4/src/entities"

func (uc *sPurchaseUsecase) UserEstimation(p *entity.UserEstimationParams) (*entity.UserEstimationResult, error) {
	items, err := uc.purchaseRepo.UserEstimation(p)

	return items, err
}
