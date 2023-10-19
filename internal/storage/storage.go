package storage

import "github.com/godofprodev/sessionauth/internal/models"

type Storage interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(ID string) (*models.User, error)
}
