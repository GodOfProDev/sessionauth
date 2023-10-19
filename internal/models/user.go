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
	Username string `validate:"required,min=4,max=15" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=7,max=20" json:"password"`
}

type LoginUserParams struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
