package purchaseRepository

import (
	"fmt"
	"log"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils"
	formatTime "projectsprintw4/src/utils/time"

	"github.com/mmcloughlin/geohash"
)

func (r *sPurchaseRepository) ListMerchantNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {

	precision := uint(3) // Adjust precision as needed
	targetGeohash := geohash.EncodeWithPrecision(filters.Lat, filters.Long, precision)
	println("targetGeohash: ", targetGeohash)
	baseQuery := `SELECT id, name, category, location_lat, location_long, created_at from merchants`
	// Print the generated SQL query and arguments
	fmt.Printf("SQL Query: %s\n", baseQuery)
	// fmt.Printf("Arguments: %v\n", args)
	rows, err := r.DB.Query(baseQuery)
	if err != nil {
		log.Printf("Error finding merchants nearby: %s", err)
		return nil, err
	}

	defer rows.Close()

	hasRows := rows.Next()
	if !hasRows {
		fmt.Println("No data found")
		return &[]entity.ListNearbyMerchantResult{}, nil
	}

	var merchants []entity.ListNearbyMerchantResult

	for rows.Next() {
		var merchant entity.ListNearbyMerchantResult
		if err := rows.Scan(
			&merchant.Id,
			&merchant.Name,
			&merchant.MerchantCategory,
			&merchant.Location.LocationLat,
			&merchant.Location.LocationLong,
			&merchant.CreatedAt); err != nil {
			log.Fatal(err)
		}

		merchant.CreatedAt, err = formatTime.FormatToISO8601WithNano(merchant.CreatedAt)

		if err != nil {
			log.Printf("Error formatting date: %s", err)
			continue
		}

		distance := utils.Haversine(filters.Lat, filters.Long, merchant.Location.LocationLat, merchant.Location.LocationLong)
		merchant.Distance = distance
		merchants = append(merchants, merchant)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &merchants, nil

}
