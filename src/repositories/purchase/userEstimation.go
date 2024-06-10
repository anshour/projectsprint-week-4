package purchaseRepository

import (
	"fmt"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"

	"github.com/lib/pq"
)

func (r *sPurchaseRepository) UserEstimation(p *entity.UserEstimationRepoParams, userId string) (*entity.UserEstimationResult, error) {

	query := `
		SELECT 
			mi.id AS item_id,
			mi.price AS item_price,
			m.id AS merchant_id,
			m.location_lat AS merchant_location_lat,
			m.location_long AS merchant_location_long,
			(acos(
				cos(radians($1)) * cos(radians(m.location_lat)) *
				cos(radians(m.location_long) - radians($2)) +
				sin(radians($1)) * sin(radians(m.location_lat))
			) * 6371) AS distance
		FROM 
			merchant_items mi
		JOIN 
			merchants m ON m.id = mi.merchant_id
		WHERE 
			m.id = ANY($3)
			AND mi.id = ANY($4)`

	var merchantItems []*entity.MerchantBindResult
	err := r.DB.Select(&merchantItems, query, p.Location.Lat, p.Location.Long, pq.Array(p.MerchantIds), pq.Array(p.ItemIds))

	if err != nil {
		return nil, err
	}

	if len(merchantItems) < len(p.ItemIds) {
		return nil, constants.ErrMissingMerchantItem
	}

	for _, item := range merchantItems {
		fmt.Printf("Distance=%f km\n", item.Distance)
		if item.Distance > 3 {
			return nil, constants.ErrTooFarLocation
		}
	}

	querySaveOrder := `INSERT INTO orders (user_id, order_status, location_lat, location_long) VALUES ($1, $2, $3, $4) RETURNING id`
	var orderId string
	if userId != "" {
		err = r.DB.QueryRowx(querySaveOrder, userId, constants.DRAFT, p.Location.Lat, p.Location.Long).Scan(&orderId)

		if err != nil {
			println("Error in saving to orders")
			return nil, err
		}
	} else {
		println(constants.ErrEmptyUserId)
		return nil, constants.ErrEmptyUserId
	}
	var totalPrice int32
	for _, item := range merchantItems {
		querySaveMerchantItem := `INSERT INTO merchant_orders (order_id, merchant_id, item_id, item_price, quantity) VALUES ($1, $2, $3, $4, $5)`
		quantity := p.ItemQuantityMap[item.ItemId]
		err := r.DB.QueryRowx(querySaveMerchantItem, orderId, item.MerchantId, item.ItemId, item.Price, quantity).Err()

		totalPrice += quantity * item.Price

		if err != nil {
			println("Error in saving to merchant_orders")
			return nil, err
		}
	}

	var estimationId string
	querySaveEstimations := `INSERT INTO estimations (order_id, total_price, estimation_minutes) VALUES ($1, $2, $3) RETURNING id`
	estimationMinutes := 0 //TODO:FIX
	err = r.DB.QueryRowx(querySaveEstimations, orderId, totalPrice, estimationMinutes).Scan(&estimationId)
	if err != nil {
		println("Error in saving to estimations")
		return nil, err
	}

	return &entity.UserEstimationResult{
		TotalPrice:         totalPrice,
		EstimationDelivery: int32(estimationMinutes),
		EstimationId:       estimationId,
	}, nil

}
