package domain

import (
	"time"

	"github.com/whalepod/milelane/app/domain/repository"
)

// DeviceType contains application platform.
type DeviceType uint

const (
	// TypeDesktop is desktop application like electron.
	TypeDesktop DeviceType = iota * 10
	// TypeIOS is iOS application.
	TypeIOS
	// TypeAndroid is Android application.
	TypeAndroid
	// TypeBrowser is distributed on web browsers.
	TypeBrowser
)

// DeviceAccessor gives access to persistence layer.
// Implementation(s) of DeviceAccessor is/are
// - DeviceRepository.
// - DeviceAccessorMock (in test).
// - DeviceAccessorErrorMock (in test).
type DeviceAccessor interface {
	Create(deviceToken string, deviceType uint) (*repository.Device, error)
}

// Device is struct for domain, not for gorm.
type Device struct {
	deviceAccessor DeviceAccessor
	UUID           string `json:"uuid"`
	DeviceToken    string `json:"device_token"`
	Type           string `json:"type"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// NewDevice returns Device struct with DeviceAccessor.
func NewDevice(da DeviceAccessor) (*Device, error) {
	var d Device
	d.deviceAccessor = da

	return &d, nil
}

// Create saves device record through persistence layer.
func (d *Device) Create(deviceToken string, deviceType string) (*Device, error) {
	repositoryDevice, err := d.deviceAccessor.Create(deviceToken, uint(GetDeviceType(deviceType)))
	if err != nil {
		return nil, err
	}

	device := Device{
		UUID:        (*repositoryDevice).UUID,
		DeviceToken: (*repositoryDevice).DeviceToken,
		Type:        DeviceType((*repositoryDevice).Type).String(),
		CreatedAt:   (*repositoryDevice).CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
		UpdatedAt:   (*repositoryDevice).UpdatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
	}

	return &device, nil
}

// String makes DeviceType compliant with Stringer interface.
func (d DeviceType) String() string {
	switch d {
	case TypeDesktop:
		return "desktop"
	case TypeIOS:
		return "ios"
	case TypeAndroid:
		return "android"
	case TypeBrowser:
		return "browser"
	default:
		return "undefined"
	}
}

// GetDeviceType returns DeviceType by string.
func GetDeviceType(deviceStr string) DeviceType {
	switch deviceStr {
	case "desktop":
		return DeviceType(TypeDesktop)
	case "ios":
		return DeviceType(TypeIOS)
	case "android":
		return DeviceType(TypeAndroid)
	case "browser":
		return DeviceType(TypeBrowser)
	default:
		return DeviceType(TypeDesktop)
	}
}
