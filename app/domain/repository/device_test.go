package repository

import (
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
