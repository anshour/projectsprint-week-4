package purchaseRepository

import (
	"fmt"
	"log"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils"
	formatTime "projectsprintw4/src/utils/time"
	"sort"
	"strconv"
	"strings"

	"github.com/mmcloughlin/geohash"
)

func (r *sPurchaseRepository) ListAllNearbyX(filters *entity.ListNearbyParams, neighbors []string) (*[]entity.ListNearbymerchantFinalResult, error) {

	baseQuery := `
	SELECT 
		merchant.id, merchant.name as merchant_name,
		merchant.category as merchant_category, merchant.location_lat, 
		merchant.location_long, merchant.image_url,
		merchant.created_at as merchant_created_at,
		merchant.geo_hash as geo_hash,
		item.id as item_id,
		item.name as item_name,  
		item.category as item_category,
		item.price as item_price, 
		item.image_url as item_image_url,
		item.created_at as item_created_at
	FROM merchants merchant 
	JOIN merchant_items item ON item.merchant_id = merchant.id
	WHERE merchant.id IN (
		SELECT id FROM merchants 
		WHERE EXISTS (
			SELECT 1 
			FROM merchant_items 
			WHERE merchant_items.merchant_id = merchants.id
		) ORDER BY id
	) AND true`

	// endQuery := fmt.Sprintf(`AND geo_hash ILIKE $1 || '%'`, "%"+targetGeohash+"%")
	conditions := []string{}
	args := []interface{}{}
	argCounter := 1

	if filters.MerchantId != "" {
		conditions = append(conditions, fmt.Sprintf("merchant.id = $%d", argCounter))
		args = append(args, filters.MerchantId)
		argCounter++
	}

	if filters.Name != "" {
		conditions = append(conditions, fmt.Sprintf("merchant.name ILIKE $%d", argCounter))
		args = append(args, "%"+filters.Name+"%")
		argCounter++
	}

	if filters.MerchantCategory != "" {
		conditions = append(conditions, fmt.Sprintf("merchant.category = $%d", argCounter))
		args = append(args, filters.MerchantCategory)
		argCounter++
	}

	precision := uint(3) // Adjust precision as needed
	targetGeohash := geohash.EncodeWithPrecision(filters.Lat, filters.Long, precision)

	baseQuery += " AND geo_hash ILIKE " + "'" + targetGeohash + "' || '%'"

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	limit := filters.Limit
	if limit == 0 {
		limit = 1000
	}
	baseQuery += fmt.Sprintf(" LIMIT $%d", argCounter)
	args = append(args, limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		baseQuery += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.Offset)
	}

	// Print the generated SQL query and arguments
	fmt.Printf("SQL Query: %s\n", baseQuery)
	// fmt.Printf("Arguments: %v\n", args)
	rows, err := r.DB.Queryx(baseQuery, args...)
	if err != nil {
		log.Printf("Error finding merchants nearby: %s", err)
		return nil, err
	}

	defer rows.Close()

	hasRows := rows.Next()
	if !hasRows {
		fmt.Println("No data found")
		return &[]entity.ListNearbymerchantFinalResult{}, nil
	}
	merchantsMap := make(map[string]*entity.ListNearbymerchantFinalResult)

	for rows.Next() {
		var merchant entity.ListNearbyMerchantResult
		var merchantItem entity.ListNearbyMerchantItemResult

		if err := rows.Scan(
			&merchant.Id,
			&merchant.Name,
			&merchant.MerchantCategory,
			&merchant.Location.LocationLat,
			&merchant.Location.LocationLong,
			&merchant.ImageUrl,
			&merchant.CreatedAt,
			&merchant.GeoHash,
			&merchantItem.Id,
			&merchantItem.Name,
			&merchantItem.Category,
			&merchantItem.Price,
			&merchantItem.ImageUrl,
			&merchantItem.CreatedAt); err != nil {
			return nil, err
		}

		merchant.CreatedAt, err = formatTime.FormatToISO8601WithNano(merchant.CreatedAt)
		if err != nil {
			log.Printf("Error formatting date: %s", err)
			continue
		}

		merchantItem.CreatedAt, err = formatTime.FormatToISO8601WithNano(merchantItem.CreatedAt)
		if err != nil {
			log.Printf("Error formatting date: %s", err)
			continue
		}

		distance := utils.Haversine(filters.Lat, filters.Long, merchant.Location.LocationLat, merchant.Location.LocationLong)
		merchant.Distance = distance

		if _, exists := merchantsMap[merchant.Id]; !exists {
			merchantsMap[merchant.Id] = &entity.ListNearbymerchantFinalResult{
				Merchant: merchant,
				Items:    []entity.ListMerchantItem{},
			}
		}

		// merchantsMap[merchant.Id].Items = append(merchantsMap[merchant.Id].Items, merchantItem)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	var merchants []entity.ListNearbymerchantFinalResult
	for _, merchant := range merchantsMap {
		fmt.Printf("looping: ID=%s, Name=%s, Distance=%f \n",
			merchant.Merchant.Id, merchant.Merchant.Name, merchant.Merchant.Distance)
		merchants = append(merchants, *merchant)
	}

	sort.Slice(merchants, func(i, j int) bool {
		return merchants[i].Merchant.Distance < merchants[j].Merchant.Distance
	})

	return &merchants, nil
}
