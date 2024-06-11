package userRepository

import (
	"log"
	entity "projectsprintw4/src/entities"
)

func (r *sUserRepository) Save(data *entity.UserSaveParam) (*entity.UserSaveResult, error) {
	query := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, username, password, email, role"

	var result entity.UserSaveResult
	err := r.DB.QueryRowx(query, data.Username, data.Email, data.Password, data.Role).StructScan(&result)
	log.Println(result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
