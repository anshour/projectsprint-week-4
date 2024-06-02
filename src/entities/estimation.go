package entity

type OrderItems struct {
	ItemId   string `json:"itemId"`
	Quantity int32  `json:"quantity"`
}

type orders struct {
	MerchantId      string       `json:"merchantId"`
	IsStartingPoint bool         `json:"isStartingPoint"`
	Items           []OrderItems `json:"items"`
}
type UserEstimationParams struct {
	Location location `json:"userLocation"`
	Orders   []orders `json:"orders"`
}

type UserEstimationRepoParams struct {
	MerchantIds     []string
	ItemIds         []string
	ItemQuantityMap map[string]int32
	Location        location
}

type UserEstimationResult struct {
	TotalPrice         int32  `json:"totalPrice"`
	EstimationDelivery int32  `json:"estimatedDeliveryTimeInMinutes"`
	EstimationId       string `json:"calculatedEstimateId"`
}
