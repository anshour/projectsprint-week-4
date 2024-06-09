package purchaseUsecase

import (
	"log"
	entity "projectsprintw4/src/entities"
)

type Point struct {
	X, Y float64
}

func (uc *sPurchaseUsecase) ListNearby(p *entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error) {

	// get merchant based on nearest location sorted and calculated by haversine
	merchants, err := uc.purchaseRepo.ListMerchantNearby(p)

	if err != nil {
		empty := make([]entity.ListNearbymerchantFinalResult, 0)
		return &empty, err
	}

	var itemIds []string
	for _, item := range *merchants {
		itemIds = append(itemIds, item.Id)
	}

	merchantItems, err := uc.purchaseRepo.GetItemMerchant(itemIds)
	if err != nil {
		// Handle the error appropriately, for example:
		log.Fatalf("Failed to get item merchant: %v", err)
	}

	merchantsMap := make(map[string]*entity.ListNearbymerchantFinalResult)
	for _, merchant := range *merchants {

		if _, exists := merchantsMap[merchant.Id]; !exists {
			merchantsMap[merchant.Id] = &entity.ListNearbymerchantFinalResult{
				Merchant: merchant,
				Items:    []entity.ListMerchantItem{},
			}
			for _, item := range *merchantItems {
				if merchant.Id == item.MerchantId {
					merchantsMap[merchant.Id].Items = append(merchantsMap[merchant.Id].Items, entity.ListMerchantItem{
						Id:        item.Id,
						Name:      item.Name,
						Category:  item.Category,
						Price:     item.Price,
						ImageUrl:  item.ImageUrl,
						CreatedAt: item.CreatedAt,
					})
				}
			}
		}

	}

	var merchantsData []entity.ListNearbymerchantFinalResult
	for _, merchant := range merchantsMap {
		merchantsData = append(merchantsData, *merchant)
	}

	// Sort locations by distance
	// sort.Slice(merchantsData, func(i, j int) bool {
	// 	return merchantsData[i].Merchant.Distance < merchantsData[j].Merchant.Distance
	// })

	// for i, merchant := range merchantsData {
	// 	if i >= p.Limit {
	// 		break
	// 	}
	// 	fmt.Printf("Location %d: ID=%s, Name=%s, Distance=%.2f km\n",
	// 		i+1, merchant.Merchant.Id, merchant.Merchant.Name, merchant.Merchant.Distance)
	// }

	return &merchantsData, nil
}
