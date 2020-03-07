package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/domain"
	"github.com/whalepod/milelane/app/domain/repository"

	"github.com/whalepod/milelane/app/infrastructure"
)

// TaskCreateJSON is struct for binding request params.
type TaskCreateJSON struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}

// TaskIndex returns all tasks.
func TaskIndex(c *gin.Context) {
	taskAccessor := repository.NewTask(infrastructure.DB)
	t, _ := domain.NewTask(taskAccessor)

	tasks, _ := t.List()

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
	}

	taskID := uint(taskIDInt)

	task, err := t.Find(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
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
	}

	taskID := uint(taskIDInt)

	err = t.Complete(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
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
	}

	taskID := uint(taskIDInt)

	err = t.Lanize(taskID)
	if err != nil {
		// In this case, possible error would be record not found.
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
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
