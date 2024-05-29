package v1routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Echo
	DB   *sqlx.DB
}

func (i *V1Routes) MountAll() {
	i.MountUser()
	i.MountMerchant()
	i.MountPurchase()
	i.MountUpload()
}
