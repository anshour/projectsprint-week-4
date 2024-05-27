package validator

import (
	"errors"
)

const (
	BEVERAGE   string = "Beverage"
	FOOD       string = "Food"
	SNACK      string = "Snack"
	CONDIMENTS string = "Condiments"
	ADDITIONS  string = "Additions"
)

func ValidateProductCategory(category string) error {
	switch category {
	case BEVERAGE,
		FOOD,
		SNACK,
		CONDIMENTS,
		ADDITIONS:
		return nil
	default:
		return errors.New("invalid category product, please check your category")
	}

}
