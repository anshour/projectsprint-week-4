package purchaseUsecase

import (
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"
)

func (uc *sPurchaseUsecase) UserEstimation(p *entity.UserEstimationParams, userId string) (*entity.UserEstimationResult, error) {
	merchantIds := []string{}
	itemIds := []string{}
	itemQuantityMap := make(map[string]int32)
	for _, order := range p.Orders {

		if order.MerchantId == "" {
			return nil, constants.ErrEmptyStringUUID
		}

		if !validator.IsValidUUID(order.MerchantId) {
			return nil, constants.ErrInvalidUUID

		}
		merchantIds = append(merchantIds, order.MerchantId)
		for _, item := range order.Items {
			if item.ItemId == "" {
				return nil, constants.ErrEmptyStringUUID
			}

			if !validator.IsValidUUID(item.ItemId) {
				return nil, constants.ErrInvalidUUID

			}
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
