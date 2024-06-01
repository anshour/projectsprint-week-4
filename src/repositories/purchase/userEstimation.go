package purchaseRepository

import (
	"log"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

func (r *sPurchaseRepository) UserEstimation(p *entity.UserEstimationRepoParams) ([]*entity.MerchantBindResult, error) {
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
			m.id IN (?)
			AND mi.id IN (?)`

	// Expand the slice into the query using sqlx.In
	query, args, err := sqlx.In(query, p.MerchantIds, p.ItemIds)
	if err != nil {
		log.Fatalln(err)
	}

	// Rebind the query for the target database (PostgreSQL in this case)
	query = r.DB.Rebind(query)

	// Execute the query
	var merchantItems []*entity.MerchantBindResult
	err = r.DB.Select(&merchantItems, query, args...)
	if err != nil {
		return nil, err
	}
	if len(merchantItems) < len(p.ItemIds) {
		return nil, constants.ErrMissingMerchantItem
	}

	return merchantItems, nil

}
