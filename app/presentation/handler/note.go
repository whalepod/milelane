package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/whalepod/milelane/app/infra/repo"
)

// NoteCreateJSON is struct for binding request params.
type NoteCreateJSON struct {
	Title string `json:"title"`
	Body  string `json:"body" binding:"required,min=1"`
}

type Note struct {
	noteAccessor repo.NoteAccessor
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewNote(no repo.NoteAccessor) *Note {
	var n Note
	n.noteAccessor = no

	return &n
}

// NoteCreate save a note.
func (t *Note) NoteCreate(c *gin.Context) {
	var n NoteCreateJSON
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	err := t.noteAccessor.Create(n.Title, n.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
