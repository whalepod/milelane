package infrastructure

import (
	"os"
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
	connectDB()

	var tasks []Task
	DB.Find(&tasks)

	t.Log("Success.")
}

func TestDBConnectionPanic(t *testing.T) {
	// To restore current connection setting.
	savedDatabaseUsername := os.Getenv("MILELANE_DATABASE_USERNAME")
	savedDatabasePassword := os.Getenv("MILELANE_DATABASE_PASSWORD")
	savedDatabaseHost := os.Getenv("MILELANE_DATABASE_HOST")
	savedDatabase := os.Getenv("MILELANE_DATABASE")

	defer func() {
		recover()

		// Restore connection setting.
		os.Setenv("MILELANE_DATABASE_USERNAME", savedDatabaseUsername)
		os.Setenv("MILELANE_DATABASE_PASSWORD", savedDatabasePassword)
		os.Setenv("MILELANE_DATABASE_HOST", savedDatabaseHost)
		os.Setenv("MILELANE_DATABASE", savedDatabase)
	}()

	os.Setenv("MILELANE_DATABASE_USERNAME", "wrongusername")
	os.Setenv("MILELANE_DATABASE_PASSWORD", "wrongpassword")
	os.Setenv("MILELANE_DATABASE_HOST", "wronghost")
	os.Setenv("MILELANE_DATABASE", "wrongdatabase")

	// If connection is wrong, it might call panic.
	connectDB()

	t.Errorf("got no error")
}
