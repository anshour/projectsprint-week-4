package entity

const ADMIN_ROLE = "ADMIN"
const USER_ROLE = "USER"

type UserLoginParams struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type UserRegisterParams struct {
	Username string `json:"username" validate:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=30"`
	Email    string `json:"email" validate:"required"`
}

type UserLoginUsecaseParams struct {
	Username string
	Password string
	Email    string
	Role     string
}

type UserSaveParam struct {
	Username string
	Password string
	Email    string
	Role     string
}

type UserSaveResult struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     string `db:"role"`
}

type UserFindResult = UserSaveResult
