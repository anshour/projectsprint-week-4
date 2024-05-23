package userRepository

import "database/sql"

func (r *sUserRepository) CheckUsernameExist(u string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := r.DB.QueryRowx(query, u).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}
