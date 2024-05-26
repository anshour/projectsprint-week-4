package merchantController

import (
	"net/http"
	entity "projectsprintw4/src/entities"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (uc *sMerchantController) ListItem(c echo.Context) error {
	filters := &entity.MerchantItemListParams{}

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

	filters.MerchantId = c.Param("merchantId")

	filters.ItemId = c.QueryParam("itemId")
	filters.Name = c.QueryParam("name")
	filters.ProductCategory = c.QueryParam("productCategory")
	filters.CreatedAt = c.QueryParam("createdAt")

	items, err := uc.merchantUsecase.ListItem(filters)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Merchant items list",
		Data:    items,
	})
}
