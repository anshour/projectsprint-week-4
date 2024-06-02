package purchaseController

import (
	"net/http"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) ListOrder(c echo.Context) error {
	filters := &entity.ListOrderParams{}

	//THIS IS THE DEFAULT FILTER, DEFINED BY SYSTEM
	filters.UserId = c.Get("userId").(string)
	filters.Status = constants.PURCHASED

	//THIS IS THE FILTER BY USER REQUEST
	filters.MerchantId = c.QueryParam("merchantId")
	filters.Category = c.QueryParam("merchantCategory")
	filters.Name = c.QueryParam("name")

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}
	orders, err := uc.purchaseUsecase.ListOrder(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Order list",
		Data:    orders,
	})
}
