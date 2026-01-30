package domain

import "time"

type UserNote struct {
	NoteID    int       `json:"note_id" validate:"required"`
	UserID    int       `json:"user_id" validate:"required"`
	Title     string    `json:"note_title" validate:"required"`
	Content   string    `json:"note_content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
