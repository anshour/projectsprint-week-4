package uploadController

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type V1Upload struct {
	DB *sqlx.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Upload interface {
	UploadImage(c echo.Context) error
}

func New(v1Upload *V1Upload) iV1Upload {
	return v1Upload
}
