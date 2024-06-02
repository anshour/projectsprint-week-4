package purchaseUsecase

import (
	entity "projectsprintw4/src/entities"
)

func (uc *sPurchaseUsecase) UserEstimation(p *entity.UserEstimationParams) (*entity.UserEstimationResult, error) {

	merchantIds := []string{}
	itemIds := []string{}
	listOrders := make(map[string]int32)
	for _, order := range p.Orders {
		// fmt.Printf("order: %s\n ", order.MerchantId)
		merchantIds = append(merchantIds, order.MerchantId)
		for _, item := range order.Items {
			listOrders[item.ItemId] = item.Quantity
			// fmt.Printf("item: %s\n ", item.ItemId)
			itemIds = append(itemIds, item.ItemId)
		}
	}
	params := &entity.UserEstimationRepoParams{
		MerchantIds: merchantIds,
		ItemIds:     itemIds,
		Location:    p.Location,
	}
	items, err := uc.purchaseRepo.UserEstimation(params)

	if err != nil {
		return nil, err
	}

	totalPrice := 0
	for _, item := range items {
		totalPrice = totalPrice + int(listOrders[item.ItemId])*int(item.Price)
	}

	estimationData := &entity.UserEstimationResult{
		TotalPrice:         totalPrice,
		EstimationDelivery: 10,
		EstimationId:       "",
	}
	return estimationData, err
}
