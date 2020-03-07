package repository

import (
	"errors"
	"time"
)

// Implements TaskRepository related with Device.

// DeviceTask is struct for gorm mapping.
type DeviceTask struct {
	ID         string    `gorm:"not null;index"`
	DeviceUUID string    `gorm:"not null;index"`
	TaskID     uint      `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"not null" sql:"type:datetime"`
	UpdatedAt  time.Time `gorm:"not null" sql:"type:datetime"`
}

// CreateDeviceTask saves Task - Device connection record.
func (t *TaskRepository) CreateDeviceTask(deviceUUID string, taskID uint) (*DeviceTask, error) {
	if deviceUUID == "" {
		return nil, errors.New("DeviceUUID can't have blank value")
	}

	deviceTask := DeviceTask{
		DeviceUUID: deviceUUID,
		TaskID:     taskID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := t.DB.Create(&deviceTask).Error; err != nil {
		return nil, err
	}

	return &deviceTask, nil
}
