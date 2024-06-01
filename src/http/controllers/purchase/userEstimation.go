package purchaseController

import (
	"net/http"
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
	data, err := uc.purchaseUsecase.UserEstimation(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Estimation",
		Data:    data,
	})
}
