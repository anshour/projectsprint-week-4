package purchaseUsecase

import (
	"fmt"
	"math"
	entity "projectsprintw4/src/entities"
	"strconv"
	"strings"

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
			fmt.Printf("km %f\n", dist)
			fmt.Printf("neearest dist %f\n", nearestDist)
			if dist < nearestDist {
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
	return &merchantNearby, nil
}
