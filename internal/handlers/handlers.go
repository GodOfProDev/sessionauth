package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/godofprodev/sessionauth/internal/session"
	"github.com/godofprodev/sessionauth/internal/storage"
)

type Handlers struct {
	store     storage.Storage
	session   session.Session
	validator *validator.Validate
}

func New(store storage.Storage, session session.Session) *Handlers {
	return &Handlers{
		store:     store,
		session:   session,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}
