package merchantUsecase

import (
	entity "projectsprintw4/src/entities"
	repository "projectsprintw4/src/repositories/merchant"
)

type sMerchantUsecase struct {
	merchantRepo repository.MerchantRepository
}

type MerchantUsecase interface {
	Create(*entity.MerchantCreateParams) (string, error)
	CreateItem(*entity.MerchantItemCreateParams) (string, error)
	List(*entity.MerchantListParams) (*[]entity.MerchantListResult, error)
}

func New(merchantRepo repository.MerchantRepository) MerchantUsecase {
	return &sMerchantUsecase{merchantRepo: merchantRepo}
}
