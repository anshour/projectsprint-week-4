package userUsecase

import (
	entity "projectsprintw4/src/entities"
	repository "projectsprintw4/src/repositories/user"
)

type sUserUsecase struct {
	userRepo repository.UserRepository
}

type UserUsecase interface {
	Login(*entity.UserLoginUsecaseParams) (string, error)
	Register(*entity.UserSaveParam) (string, error)
}

func New(userRepo repository.UserRepository) UserUsecase {
	return &sUserUsecase{userRepo: userRepo}
}
