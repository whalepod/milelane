package domain

import "github.com/whalepod/milelane/app/domain/repository"

type DeviceType uint

const (
	TypeDesktop DeviceType = iota * 10
	TypeSmartphone
	TypeTablet
	TypeIOS
	TypeAndroid
)

type DeviceAccessor interface {
	Create(deviceID string, deviceType uint) (*repository.Device, error)
}

// Struct for domain, not for gorm.
type Device struct {
	deviceAccessor DeviceAccessor
	ID             string `json:"id"`
	DeviceID       string `json:"device_id"`
	Type           string `json:"type"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func NewDevice(da DeviceAccessor) (*Device, error) {
	var d Device
	d.deviceAccessor = da

	return &d, nil
}

func (d *Device) Create(deviceID string, deviceType string) (*Device, error) {
	repositoryDevice, err := d.deviceAccessor.Create(deviceID, uint(GetDeviceType(deviceType)))
	if err != nil {
		return nil, err
	}

	device := Device{
		ID:        (*repositoryDevice).ID,
		DeviceID:  (*repositoryDevice).DeviceID,
		Type:      DeviceType((*repositoryDevice).Type).String(),
		CreatedAt: (*repositoryDevice).CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: (*repositoryDevice).UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &device, nil
}

// DeviceType is compliant with Stringer interface.
func (d DeviceType) String() string {
	switch d {
	case TypeDesktop:
		return "desktop"
	case TypeSmartphone:
		return "smartphone"
	case TypeTablet:
		return "tablet"
	case TypeIOS:
		return "ios"
	case TypeAndroid:
		return "android"
	default:
		return "undefined"
	}
}

func GetDeviceType(deviceStr string) DeviceType {
	switch deviceStr {
	case "desktop":
		return DeviceType(TypeDesktop)
	case "smartphone":
		return DeviceType(TypeSmartphone)
	case "tablet":
		return DeviceType(TypeTablet)
	case "ios":
		return DeviceType(TypeIOS)
	case "android":
		return DeviceType(TypeAndroid)
	default:
		return DeviceType(TypeDesktop)
	}
}
