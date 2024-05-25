package merchantUsecase

import entity "projectsprintw4/src/entities"

func (uc *sMerchantUsecase) List(p *entity.MerchantListParams) (*[]entity.MerchantListResult, error) {
	merchants, err := uc.merchantRepo.FindAll(p)

	if err != nil {
		empty := make([]entity.MerchantListResult, 0)
		return &empty, err
	}

	mappedMerchants := make([]entity.MerchantListResult, len(*merchants))

	for i, merchant := range *merchants {
		mappedMerchants[i].MerchantId = merchant.MerchantId
		mappedMerchants[i].Name = merchant.Name
		mappedMerchants[i].MerchantCategory = merchant.MerchantCategory
		mappedMerchants[i].ImageUrl = merchant.ImageUrl
		mappedMerchants[i].Location.Lat = merchant.LocationLat
		mappedMerchants[i].Location.Long = merchant.LocationLong
		mappedMerchants[i].CreatedAt = merchant.CreatedAt
	}

	return &mappedMerchants, nil
}
