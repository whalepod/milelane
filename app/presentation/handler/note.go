package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/infra/repo"
)

// NoteCreateJSON is struct for binding request params.
type NoteCreateJSON struct {
	Title string `json:"title" binding:"required,min=0,max=255"`
	Body  string `json:"body" binding:"required,min=1"`
}

// NoteCreate save a note.
func NoteCreate(c *gin.Context) {
	noteAccessor := repo.NewNote(infra.DB)

	var n NoteCreateJSON
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}
	fmt.Print(&n)

	err := noteAccessor.Create(n.Title, n.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
