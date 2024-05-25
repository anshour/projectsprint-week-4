package merchantRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
	"reflect"
	"strconv"
	"strings"
)

func (r *sMerchantRepository) FindAll(filters *entity.MerchantListParams) (*[]entity.MerchantFindAllResult, error) {
	baseQuery := "SELECT * FROM merchants WHERE true "

	params := []interface{}{}

	n := (&entity.MerchantListParams{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.MerchantId != "" {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantId)
		}

		if filters.Name != "" {
			conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(params)+1))
			params = append(params, "%"+filters.Name+"%")
		}

		if filters.MerchantCategory != "" {
			conditions = append(conditions, "category = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.MerchantCategory)
		}

		if len(conditions) > 0 {
			baseQuery += " AND "
			baseQuery += strings.Join(conditions, " AND ")
		}
	}

	if filters.Limit == 0 {
		filters.Limit = 5
	}

	if filters.CreatedAt == "" {
		filters.CreatedAt = "DESC"
	}

	if filters.CreatedAt == "asc" {
		filters.CreatedAt = "ASC"
	}

	if filters.CreatedAt == "desc" {
		filters.CreatedAt = "DESC"
	}

	baseQuery += " ORDER BY created_at " + filters.CreatedAt

	baseQuery += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		baseQuery += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	merchants := make([]entity.MerchantFindAllResult, filters.Limit)
	err := r.DB.Select(&merchants, baseQuery, params...)
	if err != nil {
		log.Printf("Error finding merchants: %s", err)
		return nil, err
	}

	return &merchants, nil
}
