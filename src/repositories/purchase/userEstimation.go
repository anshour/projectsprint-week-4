package purchaseRepository

import (
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
)

func (r *sPurchaseRepository) UserEstimation(p *entity.UserEstimationParams) (*entity.UserEstimationResult, error) {
	query := "INSERT INTO orders (order_status, location_lat, location_long)"

	var id string
	err := r.DB.QueryRowx(query, constants.DRAFT, p.Location.Lat, p.Location.Long).Scan(&id)

	return &entity.UserEstimationResult{}, nil

}
