package purchaseUsecase

import (
	"errors"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"
)

func (uc *sPurchaseUsecase) UserEstimation(p *entity.UserEstimationParams, userId string) (*entity.UserEstimationResult, error) {
	merchantIds := []string{}
	itemIds := []string{}
	itemQuantityMap := make(map[string]int32)

	for _, order := range p.Orders {
		if !validator.IsValidUUID(order.MerchantId) {
			println(constants.ErrInvalidType)
			return nil, errors.New(constants.ErrInvalidType)

		}
		merchantIds = append(merchantIds, order.MerchantId)
		for _, item := range order.Items {
			if !validator.IsValidUUID(item.ItemId) {
				println(constants.ErrInvalidType)
				return nil, errors.New(constants.ErrInvalidType)

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
