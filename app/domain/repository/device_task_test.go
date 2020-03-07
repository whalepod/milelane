package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	QueryDeviceTaskInsert = `INSERT INTO "device_tasks" ("device_uuid","task_id","created_at","updated_at") VALUES (?,?,?,?)`
)

func TestCreateDeviceTask(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryDeviceTaskInsert)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	taskRepository := NewTask(db)
	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	var taskID uint = 0
	_, err := taskRepository.CreateDeviceTask(deviceUUID, taskID)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateDeviceTaskWithoutDeviceUUID(t *testing.T) {
	db, _, _ := getDBMock()
	defer db.Close()

	taskRepository := NewTask(db)

	// Set blank title.
	deviceUUID := ""
	var taskID uint = 0
	_, err := taskRepository.CreateDeviceTask(deviceUUID, taskID)
	if err.Error() != "DeviceUUID can't have blank value" {
		t.Fatalf("Got %v\nwant %v", err, "DeviceUUID can't have blank value")
	}

	t.Log("Success.")
}

func TestCreateDeviceTaskRollback(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryDeviceTaskInsert)).
		WillReturnError(fmt.Errorf("Device insertion failed"))

	mock.ExpectRollback()

	taskRepository := NewTask(db)
	deviceUUID := "60982a48-9328-441f-805b-d3ab0cad9e1f"
	var taskID uint = 0
	_, err := taskRepository.CreateDeviceTask(deviceUUID, taskID)
	if err.Error() != "Device insertion failed" {
		t.Fatalf("Got %v\nwant %v", err, "Device insertion failed")
	}

	t.Log("Success.")
}
