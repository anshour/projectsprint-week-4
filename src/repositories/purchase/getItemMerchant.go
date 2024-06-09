package purchaseRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	formatTime "projectsprintw4/src/utils/time"

	"github.com/lib/pq"
)

func (r *sPurchaseRepository) GetItemMerchant(ids []string) (*[]entity.ListNearbyMerchantItemResult, error) {

	baseQuery := `
	SELECT 
		id, merchant_id, name, category, price, image_url, created_at
	FROM merchant_items 
	WHERE merchant_id = ANY($1) `

	rows, err := r.DB.Query(baseQuery, pq.Array(ids))

	if err != nil {
		log.Printf("Error finding merchant items nearby: %s", err)
		return nil, err
	}

	defer rows.Close()
	var merchantItems []entity.ListNearbyMerchantItemResult

	for rows.Next() {
		var item entity.ListNearbyMerchantItemResult

		err := rows.Scan(&item.Id, &item.MerchantId, &item.Name, &item.Category, &item.Price, &item.ImageUrl, &item.CreatedAt)
		if err != nil {
			log.Printf("Error mapping merchant items nearby: %s", err)
			return nil, err
		}

		item.CreatedAt, err = formatTime.FormatToISO8601WithNano(item.CreatedAt)
		if err != nil {
			log.Printf("Error formatting date: %s", err)
			continue
		}
		merchantItems = append(merchantItems, item)

	}

	return &merchantItems, nil

}
