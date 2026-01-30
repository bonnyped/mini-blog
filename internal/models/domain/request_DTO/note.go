package domain

type Note struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content,omitempty"`
}
