package validator

import (
	"errors"
	"regexp"
)

func ValidateUrl(url string) error {
	re := regexp.MustCompile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

	if !re.MatchString(url) {
		return errors.New("image Url is not valid")
	}

	return nil

}
