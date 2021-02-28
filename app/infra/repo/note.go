package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type NoteAccessor interface {
	Create(title string, body string) error
}

type NoteRepository struct {
	DB           *sqlx.DB
	noteAccessor NoteAccessor
}

type Note struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNote(db *sqlx.DB) NoteRepository {
	var n NoteRepository
	var na NoteAccessor

	n.DB = db
	n.noteAccessor = na

	return n
}

func (t *NoteRepository) Create(title string, body string) error {
	query := `
		INSERT INTO notes (
			title,
			body
		) VALUES (
			:title,
			:body
		);
		`

	_, err := t.DB.NamedExec(
		query,
		map[string]interface{}{
			"title": title,
			"body":  body,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
