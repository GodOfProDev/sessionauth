package storage

import (
	"github.com/godofprodev/sessionauth/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
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

const (
	createUserSQL        = `INSERT INTO users VALUES ($1, $2, $3)`
	getUserByUsernameSQL = `SELECT * FROM users WHERE username = $1`
)

func (s *PostgresStore) CreateUser(user *models.User) error {
	_, err := s.DB.Exec(createUserSQL, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetUser(username string) (*models.User, error) {
	user := new(models.User)

	err := s.DB.Get(user, getUserByUsernameSQL, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
