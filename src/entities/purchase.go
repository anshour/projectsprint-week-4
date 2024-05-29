package entity

type ListNearbyParams struct {
	LatLong          string
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
}

type ListNearbyMerchantItemResult struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	Price           int32  `json:"price"`
	ImageUrl        string `json:"imageUrl"`
	CreatedAt       string `json:"createdAt"`
}

type ListNearbyMerchantResult struct {
	MerchantId       string                       `json:"merchantId"`
	Name             string                       `json:"name"`
	MerchantCategory string                       `json:"merchantCategory"`
	ImageUrl         string                       `json:"imageUrl"`
	Location         location                     `json:"location"`
	CreatedAt        string                       `json:"createdAt"`
	Items            ListNearbyMerchantItemResult `json:"items"`
}
