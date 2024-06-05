package purchaseUsecase

import (
	"fmt"
	entity "projectsprintw4/src/entities"
	"strconv"
	"strings"

	"github.com/umahmood/haversine"
)

func (uc *sPurchaseUsecase) ListNearby(p *entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error) {
	items, err := uc.purchaseRepo.ListAllNearby(p)

	if err != nil {
		empty := make([]entity.ListNearbymerchantFinalResult, 0)
		return &empty, err
	}

	location := strings.Split(p.LatLong, ",")
	userLat, err := strconv.ParseFloat(location[0], 64)
	if err != nil {
		fmt.Println("Error parsing to float lat:", err)
		return nil, err
	}
	userLong, err := strconv.ParseFloat(location[1], 64)
	if err != nil {
		fmt.Println("Error parsing to float lat:", err)
		return nil, err
	}
	userCurrentLoc := haversine.Coord{Lat: userLat, Lon: userLong}

	merchantNearby := make([]entity.ListNearbymerchantFinalResult, 0, p.Limit)
	for _, merchant := range *items {
		merchantLoc := haversine.Coord{Lat: merchant.Merchant.Location.LocationLat, Lon: merchant.Merchant.Location.LocationLong}
		_, km := haversine.Distance(userCurrentLoc, merchantLoc)

		if km < 5 {
			merchantNearby = append(merchantNearby, merchant)
		}
	}
	return &merchantNearby, nil
}
