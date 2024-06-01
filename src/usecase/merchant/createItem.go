package merchantUsecase

import entity "projectsprintw4/src/entities"

func (uc *sMerchantUsecase) CreateItem(p *entity.MerchantItemCreateParams) (*entity.MerchantItem, error) {
	return uc.merchantRepo.SaveItem(p)
}
