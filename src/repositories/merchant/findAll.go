package merchantRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

func (r *sMerchantRepository) FindAll(filters *entity.MerchantListParams) (*[]entity.MerchantFindAllResult, error) {
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

	if filters.CreatedAt == "" || filters.CreatedAt == "desc" {
		query.AppendOrder("created_at", "DESC")
	} else if filters.CreatedAt == "asc" {
		query.AppendOrder("created_at", "ASC")
	}

	if filters.Limit == 0 {
		query.AppendLimit(5)
	} else {
		query.AppendLimit(filters.Limit)
	}

	merchants := make([]entity.MerchantFindAllResult, filters.Limit)
	err := r.DB.Select(&merchants, query.BaseQuery, query.Params...)
	if err != nil {
		log.Printf("Error finding merchants: %s", err)
		return nil, err
	}

	return &merchants, nil
}
