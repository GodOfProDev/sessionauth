package main

import (
	"github.com/godofprodev/sessionauth/internal/router"
	"github.com/godofprodev/sessionauth/internal/session"
	"github.com/godofprodev/sessionauth/internal/storage"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("there was an issue loading .env")
	}

	ps, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal("there was an issue connecting to the db")
	}

	rs, err := session.NewRedisSession()
	if err != nil {
		log.Fatal("there was an issue connecting to the db")
	}

	r := router.New(ps, rs)

	r.RegisterMiddlewares()
	r.RegisterHandlers()

	err = r.Listen()
	if err != nil {
		log.Fatal("there was an issue listening to port 8080: ", err)
	}
}
