package presentation

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, jsonStr *string) *httptest.ResponseRecorder {
	var req *http.Request

	if jsonStr != nil {
		jsonBytes := bytes.NewBuffer([]byte(*jsonStr))
		req, _ = http.NewRequest(method, path, jsonBytes)
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	return res
}

func TestRouter(t *testing.T) {
	router := Router()
	res := performRequest(router, "GET", "/", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	t.Log("Success.")
}
