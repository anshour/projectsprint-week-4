package purchaseRepository

import (
	"fmt"
	"log"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils"
	formatTime "projectsprintw4/src/utils/time"
	"strconv"
	"strings"
)

func (r *sPurchaseRepository) ListMerchantNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {

	baseQuery := `SELECT id, name, category, location_lat, location_long, created_at from merchants WHERE true`
	// Print the generated SQL query and arguments
	fmt.Printf("SQL Query: %s\n", baseQuery)
	// fmt.Printf("Arguments: %v\n", args)

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

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		baseQuery += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.Offset)
	}

	rows, err := r.DB.Queryx(baseQuery, args...)
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
