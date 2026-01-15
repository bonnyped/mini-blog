package postgres

import (
	"database/sql"
	"mini-blog/internal/logger/sl"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, sl.Err(op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(username string, creationTime time.Time) error {
	const op = "storage.postgres.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(username, created_at) VALUES($1, $2)")
	if err != nil {
		return sl.Err(op, err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(username, creationTime)

	if err != nil {
		return sl.Err(op, err)
	}

	return nil
}
