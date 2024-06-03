package validator

import (
	"errors"
	"net/mail"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid category product, please check your category")
	}
	return nil
}
