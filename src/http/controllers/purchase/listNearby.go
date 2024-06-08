package purchaseController

import (
	"fmt"
	"net/http"
	entity "projectsprintw4/src/entities"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) ListNearby(c echo.Context) error {
	filters := &entity.ListNearbyParams{}

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

	latlong := c.Param("latlong")

	if latlong == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'lat,long'",
		})
	}
	parts := strings.Split(latlong, ",")

	if len(parts) < 2 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'lat,long'",
		})
	}

	userLat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Error parsing to float lat:", err)
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'lat,long'",
		})
	}
	userLong, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'lat,long'",
		})
	}

	filters.Lat = userLat
	filters.Long = userLong
	filters.MerchantId = c.Param("merchantId")
	filters.Name = c.QueryParam("name")
	filters.MerchantCategory = c.QueryParam("merchantCategory")

	list, err := uc.purchaseUsecase.ListNearby(filters)

	if err != nil {
		println(err.Error())
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	Meta := &entity.MerchantItemMetaResult{
		Limit:  filters.Limit,
		Offset: filters.Offset,
		Total:  len(*list),
	}

	return c.JSON(http.StatusOK, SuccessResponseWithMeta{
		Message: "Merchant Nearby list",
		Data:    list,
		Meta:    Meta,
	})
}
