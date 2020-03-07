package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// DeviceRepository is struct with db connection.
type DeviceRepository struct {
	DB *gorm.DB
}

// Device is struct for gorm mapping.
type Device struct {
	UUID        string    `gorm:"not null;index"`
	DeviceToken string    `gorm:"not null;index"`
	Type        uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null" sql:"type:datetime"`
	UpdatedAt   time.Time `gorm:"not null" sql:"type:datetime"`
}

// NewDevice returns DeviceRepository with DB connection.
func NewDevice(db *gorm.DB) *DeviceRepository {
	var d DeviceRepository
	d.DB = db
	return &d
}

// Create saves device record into DB.
func (d *DeviceRepository) Create(deviceToken string, deviceType uint) (*Device, error) {
	// TODO: validate if deviceUUID is UUID or compliant to ANDROID_ID.
	if deviceToken == "" {
		return nil, errors.New("DeviceToken can't have blank value")
	}

	uuid := uuid.New().String()

	device := Device{
		UUID:        uuid,
		DeviceToken: deviceToken,
		Type:        deviceType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := d.DB.Create(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}
