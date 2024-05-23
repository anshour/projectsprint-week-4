package userUsecase

import (
	"errors"
	"log"
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/jwt"
	"projectsprintw4/src/utils/password"
)

func (u *sUserUsecase) Login(p *entity.UserLoginUsecaseParams) (token string, err error) {

	user, err := u.userRepo.FindByUsername(p.Username)
	if err != nil {
		return "", err
	}

	if user.Role != p.Role {
		return "", errors.New("role not match")
	}

	err = password.Verify(user.Password, p.Password)

	if err != nil {
		log.Printf("Error Password verify: %s", err)
		return "", constants.ErrWrongPassword
	}

	token = jwt.Generate(&jwt.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	})

	return token, nil
}
