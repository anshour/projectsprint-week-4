package userRepository

import "database/sql"

func (r *sUserRepository) CheckDataExist(u string, e string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)"
	err := r.DB.Get(&exists, query, u, e)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err != nil {
		return false, err
	}

	return exists, nil
}
