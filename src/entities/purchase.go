package entity

type ListNearbyParams struct {
	MerchantId       string
	Limit            int
	Offset           int
	Name             string
	MerchantCategory string
	Lat              float64
	Long             float64
}

type PurchaseOrderParams struct {
	EstimationId string `json:"calculatedEstimateId" validate:"required"`
}

type ListOrderParams struct {
	UserId           string
	Status           string
	MerchantId       string
	Name             string
	MerchantCategory string
	Limit            int
	Offset           int
}

type ListOrderResult struct {
	OrderId      string        `json:"orderId"`
	OrderDetails []OrderDetail `json:"orders"`
}

type OrderDetail struct {
	MerchantDetail OrderDetailMerchant `json:"merchant"`
}

type OrderDetailMerchant struct {
	Name             string                 `json:"name"`
	MerchantId       string                 `json:"merchantId"`
	MerchantCategory string                 `json:"merchantCategory"`
	ImageUrl         string                 `json:"imageUrl"`
	Location         location               `json:"location"`
	CreatedAt        string                 `json:"created_at"`
	Items            []ListOrderResultItems `json:"items"`
}

type MerchantOrderQueryResult struct {
	Id                   string
	OrderId              string
	MerchantId           string
	MerchantName         string
	MerchantCategory     string
	MerchantImageUrl     string
	MerchantLocationLat  string
	MerchantLocationLong string
	MerchantCreatedAt    string
	ItemId               string
	ItemName             string
	ItemCategory         string
	ItemPrice            int32
	ItemImageUrl         string
	ItemCreatedAt        string
	Quantity             int32
}

type ListOrderResultItems struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory" `
	ImageUrl        string `json:"imageUrl"`
	Price           int32  `json:"price"`
	Quantity        int32  `json:"quantity"`
	CreatedAt       string `json:"createdAt"`
}

type ListNearbyMerchantItemResult struct {
	Id        string `json:"itemId" db:"id"`
	Name      string `json:"name" db:"item_name"`
	Category  string `json:"productCategory" db:"item_category"`
	Price     string `json:"price" db:"item_price"`
	ImageUrl  string `json:"imageUrl" db:"item_image_url"`
	CreatedAt string `json:"createdAt" db:"item_created_at"`
}

type Location struct {
	LocationLat  float64 `json:"lat" db:"location_lat"`
	LocationLong float64 `json:"long" db:"location_long"`
}

type ListNearbyMerchantResult struct {
	Id               string `json:"merchantId" db:"id"`
	Name             string `json:"name" db:"name"`
	MerchantCategory string `json:"merchantCategory" db:"category"`
	ImageUrl         string `json:"imageUrl" db:"image_url"`
	Location         Location
	CreatedAt        string `db:"created_at"`
}

type ListNearbymerchantFinalResult struct {
	Merchant ListNearbyMerchantResult       `json:"merchant"`
	Items    []ListNearbyMerchantItemResult `json:"items"`
}

type ItemOrders struct {
	MerchantId string
	ItemId     string
	Price      string
	Quantity   int32
	OrderId    string
}
