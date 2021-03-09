package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// NoteRepository is struct with db connection.
type NoteRepository struct {
	DB *sqlx.DB
}

// NewNote returns NoteRepository with DB connection.
func NewNote(db *sqlx.DB) *NoteRepository {
	var n NoteRepository
	n.DB = db

	return &n
}

// Create saves note record into DB.
func (nr *NoteRepository) Create(title string, body string) error {
	query := `
		INSERT INTO notes (
			title,
			body,
			created_at,
			updated_at
		) VALUES (
			?,
			?,
			now(),
			now()
		);
	`

	insert, err := nr.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	insert.Exec(title, body)
	if err != nil {
		return err
	}

	return nil
}
