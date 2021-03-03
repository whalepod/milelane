package repo

import (
	"github.com/jmoiron/sqlx"
)

type NoteRepository struct {
	DB *sqlx.DB
}

func NewNote(db *sqlx.DB) *NoteRepository {
	var n NoteRepository
	n.DB = db

	return &n
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
