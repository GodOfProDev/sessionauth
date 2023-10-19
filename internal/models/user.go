package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Email    string
	Password string
}

type UserSession struct {
	ID       uuid.UUID
	Username string
	Email    string
}

type RegisterUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
