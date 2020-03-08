package presentation

import (
	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/presentation/handler"
	"github.com/whalepod/milelane/app/presentation/middleware"
)

// Router returns http router.
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSHeaders())

	r.POST("/device/create", func(c *gin.Context) {
		handler.DeviceCreate(c)
	})

	r.GET("/tasks", func(c *gin.Context) {
		handler.TaskIndex(c)
	})

	r.POST("/tasks", func(c *gin.Context) {
		handler.TaskCreate(c)
	})

	r.GET("/tasks/:taskID", func(c *gin.Context) {
		handler.TaskShow(c)
	})

	r.POST("/tasks/:taskID/complete", func(c *gin.Context) {
		handler.TaskComplete(c)
	})

	r.POST("/tasks/:taskID/lanize", func(c *gin.Context) {
		handler.TaskLanize(c)
	})

	r.POST("/tasks/:taskID/move-to-root", func(c *gin.Context) {
		handler.TaskMoveToRoot(c)
	})

	r.POST("/tasks/:taskID/move-to-child/:parentTaskID", func(c *gin.Context) {
		handler.TaskMoveToChild(c)
	})

	return r
}
