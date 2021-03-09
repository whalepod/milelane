package repo

import (
	"testing"

	"github.com/whalepod/milelane/app/infra"
)

func TestNotes_Create(t *testing.T) {
	noteRepository := NewNote(infra.DB)

	tests := []struct {
		name    string
		title   string
		body    string
		wantErr error
	}{
		{
			name:    "success",
			title:   "test title",
			body:    "test body",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := noteRepository.Create(tt.title, tt.body)
			if tt.wantErr != nil {
				if err == nil {
					t.Error("Un expected error, but got nil")
					return
				}
			}

			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
				return
			}
		})
	}
}
