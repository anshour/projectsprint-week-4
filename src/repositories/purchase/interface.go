package purchaseRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

type sPurchaseRepository struct {
	DB *sqlx.DB
}

type PurchaseRepository interface {
	UpdateOrderStatus(orderId string, status string) (err error)
	FindOrderByEstimationId(estimationId string) (orderId string, err error)
	ListAllNearby(*entity.ListNearbyParams) (*[]entity.ListNearbymerchantFinalResult, error)
	UserEstimation(*entity.UserEstimationRepoParams, string) (*entity.UserEstimationResult, error)
}

func New(DB *sqlx.DB) PurchaseRepository {
	return &sPurchaseRepository{DB: DB}
}
