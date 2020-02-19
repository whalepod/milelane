package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCORSHeaders(t *testing.T) {
	r := gin.Default()
	r.Use(CORSHeaders())
	r.GET("/", func(c *gin.Context) {})

	ts := httptest.NewServer(r)
	defer ts.Close()

	tests := []struct {
		origin string
		code   int
	}{
		// This is default localhost domain for electron.
		{"ws://127.0.0.1:5858", http.StatusOK},
	}

	for i, test := range tests {
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/", ts.URL), nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Origin", test.origin)
		res, _ := client.Do(req)
		if res == nil {
			t.Fatalf("#%d expected non-nil res", i)
		}
		defer res.Body.Close()
		if res.StatusCode != test.code {
			t.Errorf("#%d expected status code %d, got %d", i, test.code, res.StatusCode)
		}
	}
}
