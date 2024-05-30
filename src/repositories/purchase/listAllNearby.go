package purchaseRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

// SELECT
// 			item.name as item_name, item.category as item_category,
// 			item.price as item_price, merchant.category,
// 			merchant.name, merchant.location_lat, merchant.created_at, merchant.updated_at,
// 			merchant.location_long
// 			FROM merchant_items item
// 			JOIN merchants merchant ON item.merchant_id = merchant.id
// 			WHERE true;`

func (r *sPurchaseRepository) ListAllNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {
	baseQuery := `SELECT 
			merchant.id, merchant.name,
			merchant.category, merchant.location_lat, 
			merchant.location_long, merchant.image_url,
			merchant.created_at,
			item.name as item_name,  
			item.category as item_category,
			item.price as item_price, 
			item.image_url as item_image_url,
			item.created_at as item_created_at
			FROM merchant_items item 
			JOIN merchants merchant ON item.merchant_id = merchant.id
			WHERE true;`

	query := &querybuilder.Query{
		// BaseQuery: "SELECT * FROM merchants WHERE true",
		BaseQuery: baseQuery,
	}

	if filters.MerchantId != "" {
		query.AppendCondition("id", "=", filters.MerchantId)
	}

	if filters.Name != "" {
		query.AppendCondition("name", "ILIKE", "%"+filters.Name+"%")
	}

	if filters.MerchantCategory != "" {
		query.AppendCondition("category", "=", filters.MerchantCategory)
	}

	limit := filters.Limit
	if limit == 0 {
		limit = 5
	}
	// query.AppendLimit(limit)
	// merchants := make([]entity.ListNearbyMerchantResult, limit)
	rows, err := r.DB.Queryx(query.BaseQuery, query.Params...)
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
