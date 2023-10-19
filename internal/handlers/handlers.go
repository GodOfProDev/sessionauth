package handlers

import (
	"github.com/godofprodev/sessionauth/internal/session"
	"github.com/godofprodev/sessionauth/internal/storage"
	"github.com/godofprodev/sessionauth/internal/validator"
)

type Handlers struct {
	store     storage.Storage
	session   session.Session
	validator *validator.XValidator
}

func New(store storage.Storage, session session.Session) *Handlers {
	return &Handlers{
		store:     store,
		session:   session,
		validator: validator.NewValidator(),
	}
}
