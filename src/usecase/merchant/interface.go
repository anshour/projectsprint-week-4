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
	CreateItem(*entity.MerchantItemCreateParams) (*entity.MerchantItem, error)
	List(*entity.MerchantListParams) (*[]entity.MerchantListResult, error)
	ListItem(*entity.MerchantItemListParams) (*[]entity.MerchantItemListResult, error)
}

func New(merchantRepo repository.MerchantRepository) MerchantUsecase {
	return &sMerchantUsecase{merchantRepo: merchantRepo}
}
