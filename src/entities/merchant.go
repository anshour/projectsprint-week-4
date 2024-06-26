package entity

type location struct {
	Lat  float64 `json:"lat" validate:"required"`
	Long float64 `json:"long" validate:"required"`
}

type MerchantCreateParams struct {
	Name     string   `json:"name" validate:"required,min=2,max=30"`
	Category string   `json:"merchantCategory" validate:"required"` //TODO: VALIDATE ENUM
	ImageUrl string   `json:"imageUrl" validate:"required"`         //TODO: VALIDATE Image url
	Location location `json:"location" validate:"required"`
}

type MerchantItemCreateParams struct {
	Name       string `json:"name" validate:"required,min=2,max=30"`
	Category   string `json:"productCategory" validate:"required"` //TODO: VALIDATE ENUM
	Price      int32  `json:"price" validate:"required,min=1"`
	ImageUrl   string `json:"imageUrl" validate:"required"` //TODO: VALIDATE Image url
	MerchantId string
}

type MerchantListParams struct {
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
	CreatedAt        string
}

type MerchantItemListParams struct {
	MerchantId      string
	ItemId          string
	Name            string
	ProductCategory string
	CreatedAt       string
	Limit           int
	Offset          int
}

type MerchantItemListResult struct {
	Id              string `json:"itemId" db:"id"`
	Name            string `json:"name" db:"name"`
	ProductCategory string `json:"productCategory" db:"category"`
	ImageUrl        string `json:"imageUrl" db:"image_url"`
	Price           int32  `json:"price" db:"price"`
	CreatedAt       string `json:"createdAt" db:"created_at"`
}
type MerchantItemMetaResult struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type MerchantFindAllResult struct {
	MerchantId       string  `db:"id"`
	Name             string  `db:"name"`
	MerchantCategory string  `db:"category"`
	ImageUrl         string  `db:"image_url"`
	LocationLat      float64 `db:"location_lat"`
	LocationLong     float64 `db:"location_long"`
	GeoHash          string  `db:"geo_hash"`
	CreatedAt        string  `db:"created_at"`
	UpdatedAt        string  `db:"updated_at"`
}

type MerchantListResult struct {
	MerchantId       string   `json:"merchantId"`
	Name             string   `json:"name"`
	MerchantCategory string   `json:"merchantCategory"`
	ImageUrl         string   `json:"imageUrl"`
	Location         location `json:"location"`
	CreatedAt        string   `json:"createdAt"`
}

type MerchantItem struct {
	ItemId string `json:"itemId"`
}

type MerchantBindResult struct {
	MerchantId string  `db:"merchant_id"`
	Lat        float64 `db:"merchant_location_lat"`
	Long       float64 `db:"merchant_location_long"`
	ItemId     string  `db:"item_id"`
	Price      int32   `db:"item_price"`
}
