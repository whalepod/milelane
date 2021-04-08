package repo_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/infra/repo"
)

func TestNotesCreate(t *testing.T) {
	tests := []struct {
		name                string
		title               string
		body                string
		expectedError       string
		shouldInjectDBError bool
	}{
		{
			name:                "Success",
			title:               "test title",
			body:                "test body",
			expectedError:       "",
			shouldInjectDBError: false,
		},
		{
			name:                "Fail",
			title:               "test title",
			body:                "test body",
			expectedError:       "DB error",
			shouldInjectDBError: true,
		},
	}

	for _, tt := range tests {
		originalDB := infra.DB
		if tt.shouldInjectDBError {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "notes" ("title", "body", "created_at", "updated_at") VALUES (?,?,NOW(),NOW());`)).
				WithArgs(tt.title, tt.body).
				WillReturnError(fmt.Errorf("DB error"))
		}

		noteRepository := repo.NewNote(infra.DB)
		err := noteRepository.Create(tt.title, tt.body)
		if err != nil && err.Error() != tt.expectedError {
			t.Errorf("expected %s but got %s", tt.expectedError, err.Error())
		}

		infra.DB = originalDB
	}
}
