package merchantRepository

import entity "projectsprintw4/src/entities"

func (r *sMerchantRepository) SaveItem(p *entity.MerchantItemCreateParams) (*entity.MerchantItem, error) {
	query := `
	INSERT INTO merchant_items (merchant_id, name, price, category, image_url)
	SELECT m.id, $1, $2, $3, $4
	FROM merchants m
	WHERE m.id = $5 RETURNING id;
	`

	var id string
	err := r.DB.QueryRowx(query, p.Name, p.Price, p.Category, p.ImageUrl, p.MerchantId).Scan(&id)

	return &entity.MerchantItem{
		ItemId: id,
	}, err
}
