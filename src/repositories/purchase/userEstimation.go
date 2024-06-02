package purchaseRepository

import (
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"

	"github.com/lib/pq"
)

func (r *sPurchaseRepository) UserEstimation(p *entity.UserEstimationRepoParams, userId string) ([]*entity.MerchantBindResult, error) {
	query := `
		SELECT 
			m.id AS merchant_id,
			m.location_lat AS merchant_location_lat,
			m.location_long AS merchant_location_long,
			mi.id AS item_id,
			mi.price AS item_price
		FROM 
			merchants m
		JOIN 
			merchant_items mi ON m.id = mi.merchant_id
		WHERE 
			m.id = ANY($1)
			AND mi.id = ANY($2)`

	var merchantItems []*entity.MerchantBindResult
	err := r.DB.Select(&merchantItems, query, pq.Array(p.MerchantIds), pq.Array(p.ItemIds))

	if err != nil {
		return nil, err
	}

	if len(merchantItems) < len(p.ItemIds) {
		return nil, constants.ErrMissingMerchantItem
	}

	querySaveOrder := `INSERT INTO orders (user_id, order_status, location_lat, location_long) VALUES ($1, $2, $3, $4) RETURNING id`
	var orderId string

	err = r.DB.QueryRowx(querySaveOrder, userId, constants.DRAFT, p.Location.Lat, p.Location.Long).Scan(&orderId)

	if err != nil {
		println("Error in saving to orders")
		return nil, err
	}

	for _, item := range merchantItems {
		querySaveMerchantItem := `INSERT INTO merchant_orders (order_id, merchant_id, item_id, item_price, quantity) VALUES ($1, $2, $3, $4, $5)`
		//BUG: MISSING ITEM QUANTITY
		err := r.DB.QueryRowx(querySaveMerchantItem, orderId, item.MerchantId, item.ItemId, item.Price).Err()

		if err != nil {
			println("Error in saving to merchant_orders")
			return nil, err
		}
	}

	return merchantItems, nil

}
