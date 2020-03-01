package domain

import (
	"testing"

	"github.com/whalepod/milelane/app/domain/repository"
	"golang.org/x/xerrors"
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

type DeviceAccessorErrorMock struct{}

func (*DeviceAccessorErrorMock) Create(deviceID string, deviceType uint) (*repository.Device, error) {
	return nil, xerrors.New("Error mock called.")
}

var deviceTypeStrs = []string{
	"desktop",
	"smartphone",
	"tablet",
	"ios",
	"android",
	"other",
}

func TestDeviceCreate(t *testing.T) {
	var deviceAccessor DeviceAccessorMock
	device, err := NewDevice(&deviceAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceID := "dc625158-a9e9-4b7c-b15a-89991b013147"

	for _, d := range deviceTypeStrs {
		_, err = device.Create(deviceID, d)
		if err != nil {
			t.Fatalf("Returned err response: %s", err.Error())
		}
	}

	t.Log("Success.")
}

func TestDeviceCreateError(t *testing.T) {
	var deviceAccessor DeviceAccessorErrorMock
	device, err := NewDevice(&deviceAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceID := ""
	deviceType := "desktop"
	_, err = device.Create(deviceID, deviceType)
	if err.Error() != "Error mock called." {
		t.Fatalf("Got %v\nwant %v", err, "Error mock called.")
	}

	t.Log("Success: Got expected err.")
}

var deviceTypes = []struct {
	in  DeviceType
	out string
}{
	{TypeDesktop, "desktop"},
	{TypeSmartphone, "smartphone"},
	{TypeTablet, "tablet"},
	{TypeIOS, "ios"},
	{TypeAndroid, "android"},
	{999, "undefined"},
}

func TestDeviceTypeString(t *testing.T) {
	for _, d := range deviceTypes {
		deviceType := d.in
		if deviceType.String() != d.out {
			t.Fatalf("DeviceType conversion to string failed, got response: %s", deviceType.String())
		}
	}

	t.Log("Success.")
}
