package purchaseRepository

import (
	"github.com/jmoiron/sqlx"
)

type sPurchaseRepository struct {
	DB *sqlx.DB
}

type PurchaseRepository interface {
}

func New(DB *sqlx.DB) PurchaseRepository {
	return &sPurchaseRepository{DB: DB}
}
