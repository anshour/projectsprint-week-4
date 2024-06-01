package purchaseController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) UserEstimation(c echo.Context) error {

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Merchant Nearby list",
		Data:    "",
	})
}
