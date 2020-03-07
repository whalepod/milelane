package domain

import (
	"testing"
)

func TestCreateWithDevice(t *testing.T) {
	var taskAccessor TaskAccessorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	title := "Test input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateDeviceTaskError(t *testing.T) {
	var taskAccessor TaskAccessorErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	title := "Test wrong input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err.Error() != "Error mock called." {
		t.Fatalf("Got %v\nwant %v", err, "Error mock called.")
	}

	t.Log("Success: Got expected err.")
}

func TestCreateDeviceTaskErrorOnCreateDeviceTask(t *testing.T) {
	var taskAccessor TaskAccessorCreateDeviceTaskErrorMock
	task, err := NewTask(&taskAccessor)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	deviceUUID := "Wrong input."
	title := "Test input."
	_, err = task.CreateWithDevice(deviceUUID, title)
	if err.Error() != "error mock called" {
		t.Fatalf("Got %v\nwant %v", err, "error mock called")
	}

	t.Log("Success: Got expected err.")
}
