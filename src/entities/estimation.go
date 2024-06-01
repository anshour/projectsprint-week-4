package entity

type orderItems struct {
	ItemId   string `json:"itemId"`
	Quantity int32  `json:"quantity"`
}

type orders struct {
	MerchantId      string     `json:"merchantId"`
	IsStartingPoint bool       `json:"isStartingPoint"`
	Items           orderItems `json:"items"`
}
type UserEstimationParams struct {
	Location location
	Orders   orders
}

type UserEstimationResult struct {
	TotalPrice         int32  `json:"totalPrice"`
	EstimationDelivery int32  `json:"estimatedDeliveryTimeInMinutes"`
	EstimationId       string `json:"calculatedEstimateId"`
}
