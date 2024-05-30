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
	Id        string `json:"itemId" db:"id"`
	ImageUrl  string `json:"imageUrl" db:"item_image_url"`
	CreatedAt string `json:"createdAt" db:"item_created_at"`
	Name      string `json:"name" db:"item_name"`
	Price     string `json:"price" db:"item_price"`
	Category  string `json:"productCategory" db:"item_category"`
}

type Location struct {
	LocationLat  string `json:"lat" db:"location_lat"`
	LocationLong string `json:"long" db:"location_long"`
}
type ListNearbyMerchantResult struct {
	Id               string `json:"merchantId" db:"id"`
	Name             string `json:"name" db:"name"`
	MerchantCategory string `json:"merchantCategory" db:"category"`
	ImageUrl         string `json:"imageUrl" db:"image_url"`
	Location         Location
	CreatedAt        string                         `db:"created_at"`
	Items            []ListNearbyMerchantItemResult `json:"items"`
}
