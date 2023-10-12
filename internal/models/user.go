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
