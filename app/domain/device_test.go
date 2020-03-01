package domain

import (
	"testing"

	"github.com/whalepod/milelane/app/domain/repository"
)

type DeviceAccessorMock struct{}

func (*DeviceAccessorMock) Create(deviceID string, deviceType uint) (*repository.Device, error) {
	return &repository.Device{
		ID:        "60982a48-9328-441f-805b-d3ab0cad9e1f",
		DeviceID:  "dc625158-a9e9-4b7c-b15a-89991b013147",
		Type:      0,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func TestCreateDevice(t *testing.T) {
	var deviceAccessor DeviceAccessorMock
	device, err := NewDevice(&deviceAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceID := "dc625158-a9e9-4b7c-b15a-89991b013147"
	deviceType := "desktop"

	_, err = device.Create(deviceID, deviceType)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}
