package validator

import (
	"errors"
)

const (
	SMALL       string = "SmallRestaurant"
	MEDIUM      string = "MediumRestaurant"
	LARGE       string = "LargeRestaurant"
	MERCH       string = "MerchandiseRestaurant"
	BOOTH       string = "BoothKiosk"
	CONVENIENCE string = "ConvenienceStore"
)

func ValidateCategoryMerchant(category string) error {
	switch category {
	case SMALL,
		MEDIUM,
		LARGE,
		MERCH,
		BOOTH,
		CONVENIENCE:
		return nil
	default:
		return errors.New("invalid category merchant, please check your category")
	}

}
