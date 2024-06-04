package userController

import (
	"database/sql"
	"log"
	"net/http"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

func (uc *sUserController) AdminLogin(c echo.Context) error {
	var req entity.UserLoginParams

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := uc.userUsecase.Login(&entity.UserLoginUsecaseParams{
		Username: req.Username,
		Password: req.Password,
		Role:     entity.ADMIN_ROLE,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "No user found with the given username"})
		}

		if err == constants.ErrWrongPassword {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Wrong password"})
		}
		if err, ok := err.(*pgconn.PgError); ok {
			log.Println("pgsql err :", err.Error())
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
