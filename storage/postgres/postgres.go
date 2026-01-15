package postgres

import (
	"database/sql"
	"fmt"
	"mini-blog/internal/config"
	"mini-blog/internal/logger/sl"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(db_config config.DBServer) (*Storage, error) {
	const op = "storage.postgres.New"

	connStr := getPostgresConnStr(db_config)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, sl.Err(op, err)
	}

	return &Storage{db: db}, nil
}

func getPostgresConnStr(db_config config.DBServer) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		db_config.Host, db_config.Port, db_config.User, db_config.Password, db_config.Name)
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
