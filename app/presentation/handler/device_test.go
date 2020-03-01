package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/infrastructure"
)

const (
	QueryDeviceInsert = `INSERT INTO "devices" ("id","device_id","type","created_at","updated_at") VALUES (?,?,?,?,?)`
)

func TestDeviceCreate(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/device/create", func(c *gin.Context) {
		DeviceCreate(c)
	})

	// With valid title, it returns StatusOK.
	jsonStr := `{"device_id":"dc625158-a9e9-4b7c-b15a-89991b013147","device_type":"0"}`
	req, _ := http.NewRequest("POST", "/device/create", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	expectedBodyPart := "\"device_id\":\"dc625158-a9e9-4b7c-b15a-89991b013147\""
	if !strings.Contains(res.Body.String(), expectedBodyPart) {
		t.Fatalf("Returned wrong http body. Actual body: %v, Expected to have %v", res.Body.String(), expectedBodyPart)
	}

	t.Log("Success.")
}

func TestDeviceCreateWithVacantDeviceID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/device/create", func(c *gin.Context) {
		DeviceCreate(c)
	})

	// With wrong device_id, it returns StatusUnprocessableEntity.
	jsonStr := `{"device_id":""}`
	req, _ := http.NewRequest("POST", "/device/create", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusUnprocessableEntity != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestDeviceCreateWithoutDeviceID(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/device/create", func(c *gin.Context) {
		DeviceCreate(c)
	})

	// Without device_id key, it returns StatusUnprocessableEntity.
	jsonStr := `{"device_type":"0"}`
	req, _ := http.NewRequest("POST", "/device/create", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusUnprocessableEntity != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestDeviceCreateWithoutDeviceType(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/device/create", func(c *gin.Context) {
		DeviceCreate(c)
	})

	// Without device_id key, it returns StatusUnprocessableEntity.
	jsonStr := `{"device_id":"dc625158-a9e9-4b7c-b15a-89991b013147"}`
	req, _ := http.NewRequest("POST", "/device/create", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusUnprocessableEntity != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	t.Log("Success.")
}

func TestDeviceCreateFailByInfrastructure(t *testing.T) {
	res := httptest.NewRecorder()
	_, r := gin.CreateTestContext(res)
	r.POST("/device/create", func(c *gin.Context) {
		DeviceCreate(c)
	})

	// In case infrastructure.DB broken, it returns StatusInternalServerError.
	db, mock, _ := getDBMock()
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(QueryDeviceInsert)).
		WillReturnError(fmt.Errorf("Device insertion failed."))

	// Mock infrastructure.DB to test irregular error.
	originalDB := infrastructure.DB
	infrastructure.DB = db

	jsonStr := `{"device_id":"dc625158-a9e9-4b7c-b15a-89991b013147","device_type":"0"}`
	req, _ := http.NewRequest("POST", "/device/create", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(res, req)

	if http.StatusInternalServerError != res.Code {
		t.Fatalf("Returned wrong http status. Status: %v, Message: %v", res.Code, res.Body)
	}

	infrastructure.DB = originalDB
	t.Log("Success.")
}
