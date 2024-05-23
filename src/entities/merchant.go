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

// type MerchantSaveParams struct {
// 	Name         string
// 	Category     string
// 	ImageUrl     string
// 	LocationLat  float64
// 	LocationLong float64
// }

// type MerchantSaveResult struct {
// 	Id           string  `db:"id"`
// 	Name         string  `db:"name"`
// 	Category     string  `db:"category"`
// 	ImageUrl     string  `db:"image_url"`
// 	LocationLat  float64 `db:"location_lat"`
// 	LocationLong float64 `db:"location_long"`
// }
