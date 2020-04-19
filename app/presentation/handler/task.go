package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/domain"
	"github.com/whalepod/milelane/app/domain/repository"

	"github.com/whalepod/milelane/app/infrastructure"
)

// TaskCreateJSON is struct for binding request params.
type TaskCreateJSON struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

// TaskUpdateTitleJSON is struct for binding update request params.
type TaskUpdateTitleJSON struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

// TaskUpdateTermJSON is struct for binding update request params.
type TaskUpdateTermJSON struct {
	StartsAt  *time.Time `json:"starts_at" time_format:"2006-01-02T15:04:05Z"`
	ExpiresAt *time.Time `json:"expires_at" time_format:"2006-01-02T15:04:05Z"`
}

// TaskIndex returns all tasks.
func TaskIndex(c *gin.Context) {
	deviceUUID := c.GetHeader("X-Device-UUID")

	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	var tasks *[]domain.Task
	var err error
	if deviceUUID != "" {
		tasks, err = t.ListByDeviceUUID(deviceUUID)
	} else {
		tasks, err = t.List()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// TaskCreate save a task.
func TaskCreate(c *gin.Context) {
	deviceUUID := c.GetHeader("X-Device-UUID")

	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	var j TaskCreateJSON
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	var task *domain.Task
	var err error
	if deviceUUID != "" {
		task, err = t.CreateWithDevice(deviceUUID, j.Title)
	} else {
		task, err = t.Create(j.Title)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *task)
}

// TaskShow returns specific task with nested children.
func TaskShow(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	task, err := t.Find(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// TaskComplete makes a task done.
func TaskComplete(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	err = t.Complete(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// TaskUpdateTerm changes a task title.
func TaskUpdateTerm(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	var j TaskUpdateTermJSON
	// Allow request with empty body, to cleanup task term.
	// But if any wrong formatted request detected, this returns 422.
	if err := c.ShouldBindJSON(&j); err != nil && err.Error() != "invalid request" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	err = t.UpdateTerm(uint(taskIDInt), j.StartsAt, j.ExpiresAt)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// TaskUpdateTitle changes a task title.
func TaskUpdateTitle(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	var j TaskUpdateTitleJSON
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	err = t.UpdateTitle(taskID, j.Title)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// TaskLanize makes a task lanized.
func TaskLanize(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	err = t.Lanize(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// TaskMoveToRoot moves a task to root directory.
func TaskMoveToRoot(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	err = t.MoveToRoot(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// TaskMoveToChild moves a task under specified parent task.
func TaskMoveToChild(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	taskIDInt, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	taskID := uint(taskIDInt)

	parentTaskIDInt, err := strconv.Atoi(c.Param("parentTaskID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	parentTaskID := uint(parentTaskIDInt)

	err = t.MoveToChild(parentTaskID, taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
