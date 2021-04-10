package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/infra/repo"
)

// NoteCreateJSON is struct for binding request params.
type NoteCreateJSON struct {
	Title string `json:"title"`
	Body  string `json:"body" binding:"required,min=1"`
}

// Note is struct for repo.
type Note struct {
	noteAccessor repo.NoteAccessor
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// NewNote returns Note struct with NoteAccessor.
func NewNote(na repo.NoteAccessor) *Note {
	var n Note
	n.noteAccessor = na

	return &n
}

// NoteCreate save a note.
func NoteCreate(c *gin.Context) {
	noteRepo := repo.NewNote(infra.DB)
	na := NewNote(noteRepo)

	var n NoteCreateJSON
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	err := na.noteAccessor.Create(n.Title, n.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// NoteList list a note.
func NoteList(c *gin.Context) {
	noteRepo := repo.NewNote(infra.DB)
	na := NewNote(noteRepo)

	notes, err := na.noteAccessor.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}
