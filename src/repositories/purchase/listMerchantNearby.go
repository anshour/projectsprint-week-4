package purchaseRepository

import (
	"fmt"
	"log"
	entity "projectsprintw4/src/entities"
	formatTime "projectsprintw4/src/utils/time"
	"strconv"
	"strings"
)

func (r *sPurchaseRepository) ListMerchantNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {

	baseQuery := `SELECT id, name, category, location_lat, location_long, created_at,
	(acos(
		cos(radians($1)) * cos(radians(location_lat)) *
		cos(radians(location_long) - radians($2)) +
		sin(radians($1)) * sin(radians(location_lat))
		)) AS distance
	FROM merchants WHERE true`

	conditions := []string{}
	args := []interface{}{}
	argCounter := 1

	args = append(args, filters.Lat, filters.Long)

	if filters.MerchantId != "" {
		conditions = append(conditions, fmt.Sprintf("id = $%d", argCounter))
		args = append(args, filters.MerchantId)
		argCounter++
	}

	if filters.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argCounter))
		args = append(args, "%"+filters.Name+"%")
		argCounter++
	}

	if filters.MerchantCategory != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", argCounter))
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

	baseQuery += " ORDER BY distance"
	fmt.Printf("SQL Query: %s\n", baseQuery)
	fmt.Printf("Arguments: %v\n", args)
	rows, err := r.DB.Queryx(baseQuery, args...)
	if err != nil {
		log.Printf("Error finding merchants nearby: %s", err)
		return &[]entity.ListNearbyMerchantResult{}, err
	}

	defer rows.Close()

	var merchants []entity.ListNearbyMerchantResult

	for rows.Next() {
		var merchant entity.ListNearbyMerchantResult
		if err := rows.Scan(
			&merchant.Id,
			&merchant.Name,
			&merchant.MerchantCategory,
			&merchant.Location.LocationLat,
			&merchant.Location.LocationLong,
			&merchant.CreatedAt,
			&merchant.Distance); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID=%s, Name=%s, Distance=%.2f km\n",
			merchant.Id, merchant.Name, merchant.Distance)

		merchant.CreatedAt, err = formatTime.FormatToISO8601WithNano(merchant.CreatedAt)

		if err != nil {
			log.Printf("Error formatting date: %s", err)
			continue
		}

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &merchants, nil

}
