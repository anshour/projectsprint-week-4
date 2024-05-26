package merchantRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	querybuilder "projectsprintw4/src/utils/queryBuilder"
)

func (r *sMerchantRepository) FindAllItems(filters *entity.MerchantItemListParams) (*[]entity.MerchantItemListResult, error) {
	query := &querybuilder.Query{
		BaseQuery: "SELECT * FROM merchant_items WHERE true",
	}

	if filters.MerchantId != "" {
		query.AppendCondition("merchant_id", "=", filters.MerchantId)
	}

	if filters.ItemId != "" {
		query.AppendCondition("id", "=", filters.ItemId)
	}

	if filters.Name != "" {
		query.AppendCondition("name", "ILIKE", "%"+filters.Name+"%")
	}

	if filters.ProductCategory != "" {
		query.AppendCondition("category", "=", filters.ProductCategory)
	}

	if filters.CreatedAt == "" || filters.CreatedAt == "desc" {
		query.AppendOrder("created_at", "DESC")
	} else if filters.CreatedAt == "asc" {
		query.AppendOrder("created_at", "ASC")
	}

	limit := filters.Limit
	if limit == 0 {
		limit = 5
	}
	query.AppendLimit(limit)

	items := make([]entity.MerchantItemListResult, limit)
	err := r.DB.Select(&items, query.BaseQuery, query.Params...)
	if err != nil {
		log.Printf("Error finding items: %s", err)
		return nil, err
	}

	return &items, nil
}
