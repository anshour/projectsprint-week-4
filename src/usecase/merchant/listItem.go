package merchantUsecase

import entity "projectsprintw4/src/entities"

func (uc *sMerchantUsecase) ListItem(p *entity.MerchantItemListParams) (*[]entity.MerchantItemListResult, error) {
	items, err := uc.merchantRepo.FindAllItems(p)

	if err != nil {
		empty := make([]entity.MerchantItemListResult, 0)
		return &empty, err
	}

	return items, nil
}
