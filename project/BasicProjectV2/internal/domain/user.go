package domain

import "database/sql"

type User struct {
	ID       string         `json:"id"`
	Email    sql.NullString `json:"email"`
	Password string         `json:"password"`
	Phone    sql.NullString `json:"phone"`
	Birthday int            `json:"birthday"`
	Nickname string         `json:"nickname"`
	Aboutme  string         `json:"aboutme"`
}

type SignInRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
