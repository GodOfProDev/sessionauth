package handlers

import "sessionauth/internal/storage"

type Handlers struct {
	store storage.Storage
}

func New(store storage.Storage) *Handlers {
	return &Handlers{
		store: store,
	}
}
