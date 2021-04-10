package handler_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/presentation/handler"
)

func TestNoteCreate(t *testing.T) {
	tests := []struct {
		name                string
		title               string
		body                string
		DBError             bool
		wantCode            int
		shouldInjectDBError bool
	}{
		{
			name:                "Success",
			title:               "test title",
			body:                "test body",
			DBError:             false,
			wantCode:            http.StatusOK,
			shouldInjectDBError: false,
		},
		{
			name:                "Success(Empty title)",
			title:               "",
			body:                "test body",
			DBError:             false,
			wantCode:            http.StatusOK,
			shouldInjectDBError: false,
		},
		{
			name:                "Fail(binding error)",
			title:               "",
			body:                "",
			DBError:             false,
			wantCode:            http.StatusUnprocessableEntity,
			shouldInjectDBError: false,
		},
		{
			name:                "Fail(DB error)",
			title:               "",
			body:                "",
			DBError:             false,
			wantCode:            http.StatusUnprocessableEntity,
			shouldInjectDBError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			originalDB := infra.DB
			if tt.shouldInjectDBError {
				db, mock, _ := sqlmock.New()
				defer db.Close()

				mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "notes" ("title", "body", "created_at", "updated_at") VALUES (?,?,NOW(),NOW());`)).
					WillReturnError(fmt.Errorf("DB error"))
			}

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

			infra.DB = originalDB
		})
	}
}
