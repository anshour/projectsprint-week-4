package merchantRepository

import entity "projectsprintw4/src/entities"

func (r *sMerchantRepository) SaveItem(p *entity.MerchantItemCreateParams) (string, error) {
	query := "INSERT INTO merchant_items (merchant_id, name, price, category, image_url) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	var id string
	err := r.DB.QueryRowx(query, p.MerchantId, p.Name, p.Price, p.Category, p.ImageUrl).Scan(&id)

	return id, err
}
