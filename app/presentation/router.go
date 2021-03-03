package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/infra/repo"
	"github.com/whalepod/milelane/app/presentation/handler"
	"github.com/whalepod/milelane/app/presentation/middleware"
)

// Router returns http router.
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSHeaders())

	// `/` endpoint is used to healthcheck.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	noteRepository := repo.NewNote(infra.DB)
	noteHandler := handler.NewNote(noteRepository)
	r.POST("/notes", func(c *gin.Context) {
		noteHandler.NoteCreate(c)
	})

	return r
}
