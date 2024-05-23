package userRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
)

func (r *sUserRepository) FindByUsername(u string) (*entity.UserFindResult, error) {

	var result entity.UserFindResult
	err := r.DB.Get(&result, "SELECT id, username, password, email, role FROM users WHERE username = $1", u)
	if err != nil {
		log.Println("Error sql on findByUsername")
		return nil, err
	}

	return &result, nil
}
