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
	res := performRequest(router, "GET", "/tasks", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	jsonStr := `{"title":"テストタイトル"}`
	res = performRequest(router, "POST", "/tasks", &jsonStr)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	res = performRequest(router, "POST", "/tasks/1/complete", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	res = performRequest(router, "POST", "/tasks/1/lanize", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	res = performRequest(router, "POST", "/tasks/1/move-to-root", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	res = performRequest(router, "POST", "/tasks/1/move-to-child/2", nil)
	if http.StatusOK != res.Code {
		t.Fatal("Returned wrong http status.")
	}

	t.Log("Success.")
}
