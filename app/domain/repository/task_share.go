package repository

import (
	"github.com/jinzhu/gorm"
)

// TaskShareRepository is struct with DB connection.
type TaskShareRepository struct {
	DB *gorm.DB
}

// TaskShare is struct for gorm mapping.
type TaskShare struct {
	Token          uint   `gorm:"not null;index"`
	TaskID         uint   `gorm:"not null;index"`
	PermissionType string `gorm:"not null"`
}

// NewTaskShare returns TaskRepository with DB connection.
func NewTaskShare(db *gorm.DB) *TaskShareRepository {
	var t TaskShareRepository
	t.DB = db
	return &t
}

// FindByID returns list in a task share token record.
func (t *TaskShareRepository) FindByID(id uint) (*TaskShare, error) {
	var taskShare TaskShare
	if err := t.DB.Find(&taskShare, id).Error; err != nil {
		return nil, err
	}

	return &taskShare, nil
}
