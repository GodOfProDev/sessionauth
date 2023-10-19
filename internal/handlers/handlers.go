package handlers

import (
	"sessionauth/internal/session"
	"sessionauth/internal/storage"
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
