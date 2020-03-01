package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type DeviceRepository struct {
	DB *gorm.DB
}

// Struct for gorm mapping.
type Device struct {
	ID        string    `gorm:"not null;index"`
	DeviceID  string    `gorm:"not null;index"`
	Type      uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null" sql:"type:datetime"`
	UpdatedAt time.Time `gorm:"not null" sql:"type:datetime"`
}

func NewDevice(db *gorm.DB) *DeviceRepository {
	var d DeviceRepository
	d.DB = db
	return &d
}

func (d *DeviceRepository) Create(deviceID string, deviceType uint) (*Device, error) {
	// TODO: validate if deviceID is UUID or compliant to ANDROID_ID.
	if deviceID == "" {
		return nil, errors.New("DeviceID can't have blank value.")
	}

	id := uuid.New().String()

	device := Device{
		ID:        id,
		DeviceID:  deviceID,
		Type:      deviceType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := d.DB.Create(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}
