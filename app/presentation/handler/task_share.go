package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/domain"
	"github.com/whalepod/milelane/app/domain/repository"

	"github.com/whalepod/milelane/app/infrastructure"
)

// TaskShareShow returns specific task share with nested children.
func TaskShareShow(c *gin.Context) {
	taskShareAccessor := repository.NewTaskShare(infrastructure.DB)
	ts, _ := domain.NewTaskShare(taskShareAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	taskSahre, err := ts.Find(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskSahre)
}
