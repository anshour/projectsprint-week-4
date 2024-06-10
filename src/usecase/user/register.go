package userUsecase

import (
	"projectsprintw4/src/constants"
	entity "projectsprintw4/src/entities"
	"projectsprintw4/src/utils/jwt"
	"projectsprintw4/src/utils/password"
)

func (u *sUserUsecase) Register(p *entity.UserSaveParam) (token string, userId string, err error) {

	isExist, err := u.userRepo.CheckDataExist(p)

	if err != nil {
		return "", "", err
	}

	if isExist {
		return "", "", constants.ErrUsernameAlreadyExist
	}

	//TODO: VALIDATE EMAIL IS EXIST WITHIN THE SAME USER ROLE

	hashedPassword := password.Hash(p.Password)

	user, err := u.userRepo.Save(&entity.UserSaveParam{
		Username: p.Username,
		Password: hashedPassword,
		Email:    p.Email,
		Role:     p.Role,
	})

	if err != nil {
		return "", "", err
	}

	token = jwt.Generate(&jwt.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	})

	return token, user.Id, nil
}
