package purchaseRepository

func (r *sPurchaseRepository) UpdateOrderStatus(orderId string, status string) (err error) {
	query := "UPDATE orders SET order_status = $1 WHERE id = $2"
	err = r.DB.QueryRowx(query, status, orderId).Err()
	return err
}
