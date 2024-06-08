package purchaseUsecase

import (
	"math"
	entity "projectsprintw4/src/entities"

	"github.com/umahmood/haversine"
)

type Point struct {
	X, Y float64
}

func (uc *sPurchaseUsecase) ListNearby(p *entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error) {
	items, err := uc.purchaseRepo.ListAllNearby(p)
	if err != nil {
		empty := make([]entity.ListNearbymerchantFinalResult, 0)
		return &empty, err
	}

	userCurrentLoc := haversine.Coord{Lat: p.Lat, Lon: p.Long}

	nearestMerchant := make([]entity.ListNearbymerchantFinalResult, 0, p.Limit)
	var merchantNearby []entity.ListNearbymerchantFinalResult
	visited := make(map[int]bool)
	currentLoc := userCurrentLoc
	for len(*items) > len(visited) {
		nearest := -1
		nearestDist := math.MaxFloat64

		for i, merchant := range *items {
			if visited[i] {
				continue
			}
			merchantLoc := haversine.Coord{Lat: merchant.Merchant.Location.LocationLat, Lon: merchant.Merchant.Location.LocationLong}
			_, dist := haversine.Distance(currentLoc, merchantLoc)
			if dist < nearestDist {
				println("merchant: ", merchant.Merchant.Id)
				nearest = i
				nearestDist = dist
				merchantNearby = append(merchantNearby, merchant)
			}

			if nearest == -1 {
				break
			}

			visited[nearest] = true
			currentLoc = merchantLoc
		}

	}
	for i := len(merchantNearby) - 1; i >= 0; i-- {
		nearestMerchant = append(nearestMerchant, merchantNearby[i])
	}

	return &nearestMerchant, nil
}
