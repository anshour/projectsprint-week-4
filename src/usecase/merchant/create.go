package merchantUsecase

import entity "projectsprintw4/src/entities"

func (uc *sMerchantUsecase) Create(p *entity.MerchantCreateParams) (string, error) {
	return uc.merchantRepo.Save(p)
}
