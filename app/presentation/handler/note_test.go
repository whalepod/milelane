package handler_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/whalepod/milelane/app/presentation/handler"
)

func TestNoteCreate(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		body     string
		DBError  bool
		wantCode int
	}{
		{
			name:     "Success",
			title:    "test title",
			body:     "test body",
			DBError:  false,
			wantCode: http.StatusOK,
		},
		{
			name:     "Success(Empty title)",
			title:    "",
			body:     "test body",
			DBError:  false,
			wantCode: http.StatusOK,
		},
		{
			name:     "Fail(binding error)",
			title:    "",
			body:     "",
			DBError:  false,
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Create Receive
			rec := httptest.NewRecorder()
			_, r := gin.CreateTestContext(rec)
			r.POST("/notes", func(c *gin.Context) {
				handler.NoteCreate(c)
			})

			// Create Request
			body := fmt.Sprintf(`{"title": "%s", "body": "%s"}`, tt.title, tt.body)
			req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBuffer([]byte(body)))
			req.Header.Set("Content-Type", "application/json")

			// Execute
			r.ServeHTTP(rec, req)

			if diff := cmp.Diff(tt.wantCode, rec.Code); diff != "" {
				t.Errorf("mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}
