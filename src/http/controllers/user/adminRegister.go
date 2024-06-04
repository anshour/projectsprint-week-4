package userController

import (
	"log"
	"net/http"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/validator"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

func (uc *sUserController) AdminRegister(c echo.Context) error {
	var req entity.UserRegisterParams

	if err := validator.BindValidate(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}
	err := validator.ValidateEmail((req.Email))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	token, err := uc.userUsecase.Register(&entity.UserSaveParam{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     entity.ADMIN_ROLE,
	})

	if err != nil {
		if err == constants.ErrUsernameAlreadyExist {
			return c.JSON(http.StatusConflict, ErrorResponse{Message: err.Error()})
		}
		if err, ok := err.(*pgconn.PgError); ok {
			log.Println("pgsql err :", err.Error())
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"token": token,
	})
}
