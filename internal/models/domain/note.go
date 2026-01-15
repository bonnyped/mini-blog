package domain

import "time"

type Note struct {
	ID         int64     `json:"id"`
	User_id    int64     `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
