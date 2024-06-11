package purchaseController

import (
	"log"
	"net/http"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

func (uc *sPurchaseController) PurchaseOrder(c echo.Context) error {
	var req entity.PurchaseOrderParams

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	orderId, err := uc.purchaseUsecase.PurchaseOrder(req.EstimationId)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			log.Println("pgsql err :", err.Error())
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"orderId": orderId,
	})
}
