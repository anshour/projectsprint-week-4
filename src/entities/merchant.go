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

type MerchantFindAllResult struct {
	MerchantId       string  `db:"id"`
	Name             string  `db:"name"`
	MerchantCategory string  `db:"category"`
	ImageUrl         string  `db:"image_url"`
	LocationLat      float64 `db:"location_lat"`
	LocationLong     float64 `db:"location_long"`
	CreatedAt        string  `db:"created_at"`
	UpdatedAt        string  `db:"updated_at"`
}

type MerchantListResult struct {
	MerchantId       string   `json:"merchantId"`
	Name             string   `json:"name"`
	MerchantCategory string   `json:"merchantCategory"`
	ImageUrl         string   `json:"imageUrl"`
	Location         location `json:"location"`
	CreatedAt        string   `json:"created_at"`
}
