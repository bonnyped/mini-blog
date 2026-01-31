package postgres

import (
	"database/sql"
	"fmt"
	"mini-blog/internal/config"
	req "mini-blog/internal/models/domain/request_DTO"
	resp "mini-blog/internal/models/domain/responce_DTO"
	"mini-blog/pkg/sl"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type AccessToken struct {
	Token string `json:"access_token"`
}

func New(db_config config.DBServer) (*Storage, error) {
	const op = "storage.postgres.New"

	connStr := getPostgresConnStr(db_config)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, sl.Err(op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, sl.Err(op, err)
	}

	return &Storage{db: db}, nil
}

func getPostgresConnStr(db_config config.DBServer) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		db_config.Host, db_config.Port, db_config.User, db_config.Password, db_config.Name)
}

func (s *Storage) CreateUser(username string) (int, error) {
	const op = "storage.postgres.CreateUser"
	var id int

	err := s.db.QueryRow(`INSERT INTO users(username) 
						VALUES($1) RETURNING user_id`,
		username).Scan(&id)
	if err != nil {
		return -1, sl.Err(op, err)
	}

	return id, nil
}

func (s *Storage) CreateNote(userId uint64, note req.Note) error {
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

func (s *Storage) GetUserNotes(userID int) ([]resp.UserNote, error) {
	const op = "storage.postgres.GetUserNotes"
	var (
		allNotes []resp.UserNote
	)

	rows, err := s.db.Query(`SELECT note_id,
									user_id, 
									title, 
									content, 
									created_at, 
									updated_at 
							FROM notes 
							WHERE user_id=$1`, userID)
	if err != nil {
		return nil, sl.Err(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var note resp.UserNote
		err = rows.Scan(&note.NoteID,
			&note.UserID,
			&note.Title,
			&note.Content,
			&note.CreatedAt,
			&note.UpdatedAt)
		if err != nil {
			return nil, sl.Err(op, err)
		}
		allNotes = append(allNotes, note)
	}

	return allNotes, nil
}

func (s *Storage) UserExists(userID int) error {
	const op = "storage.postgres.UserExists"

	var recievedID int

	row := s.db.QueryRow(`SELECT user_id
					FROM users
					WHERE user_id=?`, userID)

	err := row.Scan(&recievedID)
	if err == sql.ErrNoRows {
		return sl.Err(op, fmt.Errorf("user with id: %d not found", userID))
	}

	return nil
}
