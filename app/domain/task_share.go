package domain

import (
	"github.com/whalepod/milelane/app/domain/repository"
)

// TaskShareAccessor is interface explaining availble methods to approach persistence layer.
type TaskShareAccessor interface {
	FindByID(id uint) (*repository.TaskShare, error)
}

// TaskShare is struct for domain, not for gorm.
type TaskShare struct {
	taskShareAccessor TaskShareAccessor
	Token             uint   `json:"token"`
	TaskID            uint   `json:"task_id"`
	PermissionType    string `json:"permission_type"`
}

// NewTaskShare returns TaskShare struct with TaskShareAccessor.
func NewTaskShare(ts TaskShareAccessor) (*TaskShare, error) {
	var t TaskShare
	t.taskShareAccessor = ts

	return &t, nil
}

// Find returns task and its descendants.
func (t *TaskShare) Find(id uint) (*TaskShare, error) {
	var taskShare TaskShare
	result, err := t.taskShareAccessor.FindByID(id)
	if err != nil {
		return &taskShare, err
	}

	return &result, nil
}
