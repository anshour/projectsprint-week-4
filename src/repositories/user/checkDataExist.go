package userRepository

import (
	"database/sql"
	entity "projectsprintw4/src/entities"
)

func (r *sUserRepository) CheckDataExist(p *entity.UserSaveParam) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR (email = $2 AND role = $3))"
	err := r.DB.Get(&exists, query, p.Username, p.Email, p.Role)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err != nil {
		return false, err
	}

	return exists, nil
}
