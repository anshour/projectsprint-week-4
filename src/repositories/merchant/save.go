package merchantRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/mmcloughlin/geohash"
)

func (r *sMerchantRepository) Save(p *entity.MerchantCreateParams) (string, error) {
	geoHash := geohash.Encode(p.Location.Lat, p.Location.Long)
	query := "INSERT INTO merchants (name, category, image_url, location_lat, location_long, geo_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	var id string
	err := r.DB.QueryRowx(query, p.Name, p.Category, p.ImageUrl, p.Location.Lat, p.Location.Long, geoHash).Scan(&id)

	return id, err
}
