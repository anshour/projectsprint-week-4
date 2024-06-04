package merchantController

import (
	"log"
	"net/http"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/labstack/echo/v4"
)

func (uc *sMerchantController) CreateItem(c echo.Context) error {
	var req entity.MerchantItemCreateParams

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	if err := validator.ValidateProductCategory(req.Category); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	if err := validator.ValidateUrl(req.ImageUrl); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	merchantId := c.Param("merchantId")

	if merchantId == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'merchant id'",
		})
	}
	itemId, err := uc.merchantUsecase.CreateItem(&entity.MerchantItemCreateParams{
		Name:       req.Name,
		Category:   req.Category,
		Price:      req.Price,
		ImageUrl:   req.ImageUrl,
		MerchantId: merchantId,
	})

	if err != nil {
		log.Println(err.Error())
		if err.Error() == constants.ErrNoRowsResultText {
			return c.JSON(http.StatusNotFound, ErrorResponse{Message: constants.ErrNoRowsResultText})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"itemId": itemId,
	})
}
