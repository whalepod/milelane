package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	QueryDeviceInsert = `INSERT INTO "devices" ("id","device_id","type","created_at","updated_at") VALUES (?,?,?,?,?)`
)

func TestCreateDevice(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryDeviceInsert)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	deviceRepository := NewDevice(db)
	deviceID := "561c36de-695d-49e8-b124-57e1f742c90a"
	var deviceType uint = 0
	_, err := deviceRepository.Create(deviceID, deviceType)
	if err != nil {
		t.Fatalf("Returned err response: %s", err.Error())
	}

	t.Log("Success.")
}

func TestCreateDeviceWithoutDeviceID(t *testing.T) {
	db, _, _ := getDBMock()
	defer db.Close()

	deviceRepository := NewDevice(db)

	// Set blank title.
	deviceID := ""
	var deviceType uint = 0
	_, err := deviceRepository.Create(deviceID, deviceType)
	if err.Error() != "DeviceID can't have blank value." {
		t.Fatalf("Got %v\nwant %v", err, "DeviceID can't have blank value.")
	}

	t.Log("Success.")
}

func TestCreateDeviceRollback(t *testing.T) {
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(QueryDeviceInsert)).
		WillReturnError(fmt.Errorf("Device insertion failed."))

	mock.ExpectRollback()

	deviceRepository := NewDevice(db)
	deviceID := "561c36de-695d-49e8-b124-57e1f742c90a"
	var deviceType uint = 0
	_, err := deviceRepository.Create(deviceID, deviceType)
	if err.Error() != "Device insertion failed." {
		t.Fatalf("Got %v\nwant %v", err, "Device insertion failed.")
	}

	t.Log("Success.")
}
