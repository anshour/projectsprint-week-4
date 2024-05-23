package userRepository

import (
	entity "projectsprintw4/src/entities"

	"github.com/jmoiron/sqlx"
)

type sUserRepository struct {
	DB *sqlx.DB
}

type UserRepository interface {
	Save(*entity.UserSaveParam) (*entity.UserSaveResult, error)
	FindByUsername(u string) (*entity.UserFindResult, error)
	CheckUsernameExist(u string) (bool, error)
}

func New(DB *sqlx.DB) UserRepository {
	return &sUserRepository{DB: DB}
}
