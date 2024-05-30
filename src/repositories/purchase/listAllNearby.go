package purchaseRepository

import (
	"fmt"
	"log"
	entity "projectsprintw4/src/entities"
	"strings"
)

func (r *sPurchaseRepository) ListAllNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {
	baseQuery := `
    SELECT 
        merchant.id, merchant.name as merchant_name,
        merchant.category as merchant_category, merchant.location_lat, 
        merchant.location_long, merchant.image_url,
        merchant.created_at as merchant_created_at,
        item.name as item_name,  
        item.category as item_category,
        item.price as item_price, 
        item.image_url as item_image_url,
        item.created_at as item_created_at
    FROM merchant_items item 
    JOIN merchants merchant ON item.merchant_id = merchant.id
    WHERE true`

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

	limit := filters.Limit
	if limit == 0 {
		limit = 5
	}
	baseQuery += fmt.Sprintf(" LIMIT $%d", argCounter)
	args = append(args, limit)

	// Print the generated SQL query and arguments
	fmt.Printf("SQL Query: %s\n", baseQuery)
	fmt.Printf("Arguments: %v\n", args)
	rows, err := r.DB.Queryx(baseQuery, args...)
	if err != nil {
		log.Printf("Error finding merchants nearby: %s", err)
		return nil, err
	}

	defer rows.Close()

	merchantsMap := make(map[string]*entity.ListNearbyMerchantResult)

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
			&merchantItem.Name,
			&merchantItem.Category,
			&merchantItem.Price,
			&merchantItem.ImageUrl,
			&merchantItem.CreatedAt); err != nil {
			return nil, err
		}

		if _, exists := merchantsMap[merchant.Id]; !exists {
			merchantsMap[merchant.Id] = &merchant
		}

		if merchantItem.Name != "" {
			merchantsMap[merchant.Id].Items = append(merchantsMap[merchant.Id].Items, merchantItem)
		}
	}

	var merchants []entity.ListNearbyMerchantResult
	for _, merchant := range merchantsMap {
		merchants = append(merchants, *merchant)
	}

	return &merchants, nil

}
