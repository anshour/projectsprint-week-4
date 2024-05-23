package validator

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func BindValidate(c echo.Context, req interface{}) (err error) {
	if err = c.Bind(req); err != nil {
		err = fmt.Errorf("failed to parse the request, err: %s", err.Error())
		return
	}

	if err = c.Validate(req); err != nil {
		err = fmt.Errorf("failed to validate the request, err: %s", err.Error())
		return
	}

	return
}
