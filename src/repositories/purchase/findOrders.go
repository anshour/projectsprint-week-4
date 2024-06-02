package purchaseRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

func (r *sPurchaseRepository) FindOrders(filters *entity.ListOrderParams) (*[]entity.ListOrderResult, error) {
	query := &querybuilder.Query{
		BaseQuery: `SELECT id FROM orders WHERE true`,
	}

	if filters.UserId != "" {
		query.AppendCondition("user_id", "=", filters.UserId)
	}

	if filters.Status != "" {
		query.AppendCondition("order_status", "=", filters.Status)
	}

	limit := filters.Limit
	if limit == 0 {
		limit = 5
	}
	query.AppendLimit(limit)

	orderIds := make([]string, 0, limit)
	rows, err := r.DB.Queryx(query.BaseQuery, query.Params...)
	if err != nil {
		log.Printf("Error finding orders: %s", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var orderId string

		if err := rows.Scan(
			&orderId,
		); err != nil {
			return nil, err
		}
		orderIds = append(orderIds, orderId)
	}

	merchantOrderQuery := &querybuilder.Query{
		BaseQuery: `SELECT 
			mo.id,
			mo.order_id,
			mo.merchant_id, 
			m.name as merchant_name, 
			m.category as merchant_category, 
			m.image_url as merchant_image_url, 
			m.location_lat as merchant_location_lat, 
			m.location_long as merchant_location_long, 
			m.created_at as merchant_created_at, 
			mo.item_id, 
			mi.name as item_name,
			mi.category as item_category,
			mi.price as item_price,
			mi.image_url as item_image_url,
			mi.created_at as item_created_at,
			mo.quantity as quantity
		FROM merchant_orders mo
		JOIN merchants m ON m.id = mo.merchant_id
		JOIN merchant_items mi ON mi.id = mo.item_id 
		WHERE true`,
	}

	merchantOrderQuery.AppendWhereAny("mo.order_id", orderIds)

	if filters.MerchantId != "" {
		merchantOrderQuery.AppendCondition("mo.merchant_id", "=", filters.MerchantId)
	}

	if filters.Name != "" {
		merchantOrderQuery.AppendOrWhere("m.name", "ILIKE", "%"+filters.Name+"%", "mi.name", "ILIKE", "%"+filters.Name+"%")
	}

	if filters.MerchantCategory != "" {
		merchantOrderQuery.AppendCondition("m.category", "=", filters.MerchantCategory)
	}

	rows, err = r.DB.Queryx(merchantOrderQuery.BaseQuery, merchantOrderQuery.Params...)
	if err != nil {
		log.Printf("Error finding order detail (merchant_orders): %s", err)
		return nil, err
	}

	defer rows.Close()

	merchantsByOrderId := make(map[string][]entity.OrderDetailMerchant, 0)
	itemsByOrderIdMerchantId := make(map[string]map[string][]entity.ListOrderResultItems, 0)

	for rows.Next() {
		var merchantOrder entity.MerchantOrderQueryResult

		if err := rows.Scan(
			&merchantOrder.Id,
			&merchantOrder.OrderId,
			&merchantOrder.MerchantId,
			&merchantOrder.MerchantName,
			&merchantOrder.MerchantCategory,
			&merchantOrder.MerchantImageUrl,
			&merchantOrder.MerchantLocationLat,
			&merchantOrder.MerchantLocationLong,
			&merchantOrder.MerchantCreatedAt,
			&merchantOrder.ItemId,
			&merchantOrder.ItemName,
			&merchantOrder.ItemCategory,
			&merchantOrder.ItemPrice,
			&merchantOrder.ItemImageUrl,
			&merchantOrder.ItemCreatedAt,
			&merchantOrder.Quantity,
		); err != nil {
			return nil, err
		}

		isMerchantIdExist := false
		if _, ok := merchantsByOrderId[merchantOrder.OrderId]; ok {
			for _, x := range merchantsByOrderId[merchantOrder.OrderId] {
				if x.MerchantId == merchantOrder.MerchantId {
					isMerchantIdExist = true
				}
			}
		}

		if !isMerchantIdExist {
			merchantsByOrderId[merchantOrder.OrderId] = append(merchantsByOrderId[merchantOrder.OrderId], entity.OrderDetailMerchant{
				Name:             merchantOrder.MerchantName,
				MerchantId:       merchantOrder.MerchantId,
				MerchantCategory: merchantOrder.MerchantCategory,
				ImageUrl:         merchantOrder.MerchantImageUrl,
				CreatedAt:        merchantOrder.MerchantCreatedAt,
			})
		}

		if _, ok := itemsByOrderIdMerchantId[merchantOrder.OrderId]; !ok {
			itemsByOrderIdMerchantId[merchantOrder.OrderId] = make(map[string][]entity.ListOrderResultItems)
		}

		itemsByOrderIdMerchantId[merchantOrder.OrderId][merchantOrder.MerchantId] = append(itemsByOrderIdMerchantId[merchantOrder.OrderId][merchantOrder.MerchantId], entity.ListOrderResultItems{
			ItemId:          merchantOrder.ItemId,
			Name:            merchantOrder.ItemName,
			ProductCategory: merchantOrder.ItemCategory,
			ImageUrl:        merchantOrder.ItemImageUrl,
			Price:           merchantOrder.ItemPrice,
			Quantity:        merchantOrder.Quantity,
			CreatedAt:       merchantOrder.ItemCreatedAt,
		})
	}

	orders := make([]entity.ListOrderResult, 0, limit)

	for _, orderId := range orderIds {
		orderDetails := make([]entity.OrderDetail, 0)

		for _, merchant := range merchantsByOrderId[orderId] {
			merchant.Items = itemsByOrderIdMerchantId[orderId][merchant.MerchantId]
			orderDetails = append(orderDetails, entity.OrderDetail{
				MerchantDetail: merchant,
			})
		}
		orders = append(orders, entity.ListOrderResult{
			OrderId:      orderId,
			OrderDetails: orderDetails,
		})
	}

	return &orders, nil

}
