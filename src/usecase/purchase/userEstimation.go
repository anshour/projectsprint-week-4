package purchaseUsecase

import (
	entity "projectsprintw4/src/entities"
)

func (uc *sPurchaseUsecase) UserEstimation(p *entity.UserEstimationParams, userId string) (*entity.UserEstimationResult, error) {
	merchantIds := []string{}
	itemIds := []string{}
	itemQuantityMap := make(map[string]int32)

	for _, order := range p.Orders {
		merchantIds = append(merchantIds, order.MerchantId)
		for _, item := range order.Items {
			itemIds = append(itemIds, item.ItemId)
			itemQuantityMap[item.ItemId] = item.Quantity
		}
	}

	params := &entity.UserEstimationRepoParams{
		MerchantIds:     merchantIds,
		ItemIds:         itemIds,
		Location:        p.Location,
		ItemQuantityMap: itemQuantityMap,
	}

	return uc.purchaseRepo.UserEstimation(params, userId)
}
