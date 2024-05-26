package merchantController

import (
	"log"
	"net/http"
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
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Merchant item created",
		Data: map[string]any{
			"itemId": itemId,
		},
	})
}
