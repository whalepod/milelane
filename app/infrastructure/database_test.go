package infrastructure

import (
	"testing"
)

// Struct for gorm mapping.
type Task struct {
	ID          uint   `gorm:"not null;index"`
	Title       string `gorm:"not null;index"`
	CompletedAt string
	CreatedAt   string `gorm:"not null"`
	UpdatedAt   string `gorm:"not null"`
}

func TestDB(t *testing.T) {
	// Test query with core table name - tasks.
	// If connection is wrong, it might call panic.
	var tasks []Task
	DB.Find(&tasks)

	t.Log("Success.")
}
