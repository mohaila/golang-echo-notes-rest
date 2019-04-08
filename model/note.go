package model

type (
	// Note models a note
	Note struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)
