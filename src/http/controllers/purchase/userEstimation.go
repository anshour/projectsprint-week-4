package purchaseController

import (
	"net/http"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) UserEstimation(c echo.Context) error {
	var req entity.UserEstimationParams
	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	if len(req.Orders) == 0 {
		return c.JSON(http.StatusNotFound, ErrorResponse{
			Message: "Empty order",
			Status:  false,
		})
	}
	UserId, _ := c.Get("userId").(string)
	data, err := uc.purchaseUsecase.UserEstimation(&req, UserId)

	if err != nil {
		println("err.Error(): ", err.Error())
		if err.Error() == constants.ErrInvalidType {
			return c.JSON(http.StatusNotFound, ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"totalPrice":                     data.TotalPrice,
		"estimatedDeliveryTimeInMinutes": data.EstimationDelivery,
		"calculatedEstimateId":           data.EstimationId,
	})
}
