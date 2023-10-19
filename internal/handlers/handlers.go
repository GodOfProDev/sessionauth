package handlers

import (
	"github.com/godofprodev/sessionauth/internal/session"
	"github.com/godofprodev/sessionauth/internal/storage"
)

type Handlers struct {
	store   storage.Storage
	session session.Session
}

func New(store storage.Storage, session session.Session) *Handlers {
	return &Handlers{
		store:   store,
		session: session,
	}
}
