package purchaseRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

func (r *sPurchaseRepository) ListAllNearby(filters *entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error) {
	query := &querybuilder.Query{
		BaseQuery: "SELECT * FROM merchants WHERE true",
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
	query.AppendLimit(limit)

	merchants := make([]entity.ListNearbyMerchantResult, limit)
	err := r.DB.Select(&merchants, query.BaseQuery, query.Params...)
	if err != nil {
		log.Printf("Error finding merchants nearby: %s", err)
		return nil, err
	}

	return &merchants, nil

}
