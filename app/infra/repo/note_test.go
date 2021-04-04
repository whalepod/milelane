package repo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/whalepod/milelane/app/infra"
	"github.com/whalepod/milelane/app/infra/repo"
)

func TestNotesCreate(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		body    string
		DBError bool
		wantErr string
	}{
		{
			name:    "Success",
			title:   "test title",
			body:    "test body",
			DBError: false,
			wantErr: "",
		},
		{
			name:    "Fail",
			title:   "test title",
			body:    "test body",
			DBError: true,
			wantErr: "DB error",
		},
	}

	noteRepository := repo.NewNote(infra.DB)
	for _, tt := range tests {
		originalDB := infra.DB

		t.Run(tt.name, func(t *testing.T) {
			if tt.DBError {
				db, mock, _ := sqlmock.New()
				defer db.Close()

				mock.ExpectPrepare("INSERT INTO").
					WillBeClosed().
					ExpectExec().
					WithArgs(tt.title, tt.body, time.Now(), time.Now()).
					WillReturnError(fmt.Errorf("DB error"))

				infra.DB = sqlx.NewDb(db, "sqlmock")
			}

			noteRepository = repo.NewNote(infra.DB)
			err := noteRepository.Create(tt.title, tt.body)
			if err != nil && err.Error() != tt.wantErr {
				t.Errorf("expected %s but got %s", tt.wantErr, err.Error())
			}

			infra.DB = originalDB
		})
	}
}
