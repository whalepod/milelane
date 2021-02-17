package infrastructure

import (
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	connectDB()

	// TODO
	// Find some data from DB.
	// For now, this test does nothing.

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
