package merchantRepository

import entity "projectsprintw4/src/entities"

func (r *sMerchantRepository) Save(p *entity.MerchantCreateParams) (string, error) {
	query := "INSERT INTO merchants (name, category, image_url, location_lat, location_long) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id string
	err := r.DB.QueryRowx(query, p.Name, p.Category, p.ImageUrl, p.Location.Lat, p.Location.Long).Scan(&id)

	return id, err
}
