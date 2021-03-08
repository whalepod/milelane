package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/whalepod/milelane/app/infra/repo"
)

type mockNoteAccesor struct {
	mock.Mock
	repo.NoteAccessor
}

func (m *mockNoteAccesor) NoteCreate(ctx *gin.Context) error {
	return m.Called(ctx).Error(0)
}

func TestNotes_NoteCreate(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		body     string
		mockErr  error
		wantCode int
		wantBody error
	}{
		{
			name:     "Success",
			title:    "test title",
			body:     "test body",
			mockErr:  nil,
			wantCode: http.StatusOK,
			wantBody: nil,
		},
		{
			name:     "Success(Empty title)",
			title:    "",
			body:     "test body",
			mockErr:  nil,
			wantCode: http.StatusOK,
			wantBody: nil,
		},
		{
			name:     "body required",
			title:    "",
			body:     "",
			mockErr:  nil,
			wantCode: http.StatusInternalServerError,
			wantBody: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create Mock
			mock := &mockNoteAccesor{}
			mock.On("Create").Return(tt.mockErr)

			// Create Receive
			res := httptest.NewRecorder()
			_, r := gin.CreateTestContext(res)
			h := NewNote(mock)
			r.POST("/notes", func(c *gin.Context) {
				h.NoteCreate(c)
			})

			// Create Request
			body := fmt.Sprintf(`{"title": "%s", "body": "%s"}`, tt.title, tt.body)
			req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBuffer([]byte(body)))
			req.Header.Set("Content-Type", "application/json")

			// Execute
			r.ServeHTTP(res, req)

			if diff := cmp.Diff(tt.wantBody, req.Body); diff != "" {
				t.Errorf("mismatch body (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantCode, req.Response.StatusCode); diff != "" {
				t.Errorf("mismatch code (-want +got):\n%s", diff)
			}
		})
	}
}
