package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"sessionauth/internal/models"
)

type PostgresStore struct {
	DB *sqlx.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	url := os.Getenv("POSTGRES_DB_URL")
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (s *PostgresStore) CreateUser(*models.User) error {
	return nil
}

func (s *PostgresStore) GetUser(username string) (*models.User, error) {
	return nil, nil
}
