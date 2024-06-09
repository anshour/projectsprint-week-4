package purchaseRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

type sPurchaseRepository struct {
	DB *sqlx.DB
}

type PurchaseRepository interface {
	FindOrders(*entity.ListOrderParams) (*[]entity.ListOrderResult, error)
	UpdateOrderStatus(orderId string, status string) (err error)
	GetOrderIdByEstimationId(estimationId string) (orderId string, err error)
	ListMerchantNearby(*entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error)
	GetItemMerchant([]string) (*[]entity.ListNearbyMerchantItemResult, error)
	UserEstimation(*entity.UserEstimationRepoParams, string) (*entity.UserEstimationResult, error)
}

func New(DB *sqlx.DB) PurchaseRepository {
	return &sPurchaseRepository{DB: DB}
}
