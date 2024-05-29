package purchaseRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

type sPurchaseRepository struct {
	DB *sqlx.DB
}

type PurchaseRepository interface {
	ListAllNearby(*entity.ListNearbyParams) (*[]entity.ListNearbyMerchantResult, error)
}

func New(DB *sqlx.DB) PurchaseRepository {
	return &sPurchaseRepository{DB: DB}
}
