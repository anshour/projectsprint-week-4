package purchaseRepository

func (r *sPurchaseRepository) GetOrderIdByEstimationId(estimationId string) (orderId string, err error) {
	query := "SELECT order_id FROM estimations WHERE id = $1"
	err = r.DB.QueryRowx(query, estimationId).Scan(&orderId)
	return orderId, err
}
