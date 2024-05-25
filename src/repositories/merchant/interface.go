package merchantRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

type sMerchantRepository struct {
	DB *sqlx.DB
}

type MerchantRepository interface {
	Save(*entity.MerchantCreateParams) (string, error)
	FindAll(*entity.MerchantListParams) (*[]entity.MerchantFindAllResult, error)
}

func New(DB *sqlx.DB) MerchantRepository {
	return &sMerchantRepository{DB: DB}
}
