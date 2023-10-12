package main

import (
	"github.com/joho/godotenv"
	"log"
	"sessionauth/internal/router"
	"sessionauth/internal/session"
	"sessionauth/internal/storage"
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

	rs, err := session.NewRedisSession(ps)
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
