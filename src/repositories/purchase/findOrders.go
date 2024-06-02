package purchaseRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

func (r *sPurchaseRepository) FindOrders(filters *entity.ListOrderParams) (*[]entity.ListOrderResult, error) {
	query := &querybuilder.Query{
		BaseQuery: `SELECT o.id
		FROM orders o
		 WHERE true`,
	}

	if filters.UserId != "" {
		query.AppendCondition("user_id", "=", filters.UserId)
	}

	if filters.Status != "" {
		query.AppendCondition("order_status", "=", filters.Status)
	}

	if filters.MerchantId != "" {
		//TODO: HANDlE
		// query.AppendCondition("id", "=", filters.MerchantId)
	}

	if filters.Name != "" {
		//TODO: HANDlE
		// query.AppendCondition("name", "ILIKE", "%"+filters.Name+"%")
	}

	if filters.Category != "" {
		//TODO: HANDlE
		// query.AppendCondition("category", "=", filters.MerchantCategory)
	}

	limit := filters.Limit
	if limit == 0 {
		limit = 5
	}
	query.AppendLimit(limit)

	orders := make([]entity.ListOrderResult, 0, limit)
	log.Println(query)
	rows, err := r.DB.Queryx(query.BaseQuery, query.Params...)
	if err != nil {
		log.Printf("Error finding orders: %s", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var order entity.ListOrderResult

		if err := rows.Scan(
			&order.OrderId,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return &orders, nil

}
