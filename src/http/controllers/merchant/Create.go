package merchantController

import (
	"log"
	"net/http"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/labstack/echo/v4"
)

func (uc *sMerchantController) Create(c echo.Context) error {
	var req entity.MerchantCreateParams

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	merchantId, err := uc.merchantUsecase.Create(&req)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Merchant created",
		Data: map[string]any{
			"merchantId": merchantId,
		},
	})
}
