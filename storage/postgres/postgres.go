package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"mini-blog/internal/config"
	"mini-blog/internal/logger/sl"
	"mini-blog/internal/models/domain"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

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

func (s *Storage) CreateUser(logger slog.Logger, username string, secret string) error {
	const op = "storage.postgres.CreateUser"
	var id int64

	err := s.db.QueryRow("INSERT INTO users(username) VALUES($1) RETURNING id", username).Scan(&id)
	if err != nil {
		logger.Error("NEED TO DISCRIBE")
		return sl.Err(op, err)
	}

	resJWT, err := generateToken(id, secret)
	if err != nil {
		logger.Error("NEED TO DISCRIBE4")
		return sl.Err(op, err)
	}

	logger.Info("Generated JWT token")
	logger.Info("JWT: ", "token", resJWT)

	return nil
}

func generateToken(id int64, secret string) (string, error) {
	key := []byte(secret)
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(id, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resJWT, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return resJWT, nil
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
