package storage

import "sessionauth/internal/models"

type Storage interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
}
