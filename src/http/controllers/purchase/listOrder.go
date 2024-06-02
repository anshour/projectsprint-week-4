package purchaseController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) ListOrder(c echo.Context) error {

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Order list",
		Data:    nil,
	})
}
