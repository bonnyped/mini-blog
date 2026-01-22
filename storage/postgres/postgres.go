package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"mini-blog/internal/config"
	"mini-blog/internal/logger/sl"
	"mini-blog/internal/models/domain"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(logger *slog.Logger, db_config config.DBServer) (*Storage, error) {
	const op = "storage.postgres.New"

	connStr := getPostgresConnStr(db_config)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, sl.Err(op, err)
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Failed to connect to Postgres", "op", op, "error", err)
		return nil, sl.Err(op, err)
	}
	logger.Info("Connected to Postgres successfully")

	return &Storage{db: db}, nil
}

func getPostgresConnStr(db_config config.DBServer) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		db_config.Host, db_config.Port, db_config.User, db_config.Password, db_config.Name)
}

func (s *Storage) CreateUser(username string) error {
	const op = "storage.postgres.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(username) VALUES($1)")
	if err != nil {
		return sl.Err(op, err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(username)

	if err != nil {
		return sl.Err(op, err)
	}

	return nil
}

func (s *Storage) CreateNote(userId uint64, note domain.Note) error {
	const op = "storage.posgtres.CreateNote"

	stmt, err := s.db.Prepare("INSERT INTO notes(user_id, title, content) VALUES($1, $2, $3)")
	if err != nil {
		return sl.Err(op, err)
	}

	_, err = stmt.Exec(userId, note.Title, note.Content)
	if err != nil {
		return sl.Err(op, err)
	}

	return nil
}
