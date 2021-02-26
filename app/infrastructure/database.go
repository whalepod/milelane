package infrastructure

import (
	"os"
	"time"

	// To enable sqlx to connect MySQL.
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB exposes database connection.
var DB *sqlx.DB

const maxSleepTime = 10

type mySQLConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

func init() {
	DB = connectDB()
}

func connectDB() *sqlx.DB {
	// Initalize db connection.
	m := mySQLConfig{
		Username: os.Getenv("MILELANE_DATABASE_USERNAME"),
		Password: os.Getenv("MILELANE_DATABASE_PASSWORD"),
		Host:     os.Getenv("MILELANE_DATABASE_HOST"),
		Database: os.Getenv("MILELANE_DATABASE"),
	}

	dbConfigStr := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":3306)/" + m.Database + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dbConfigStr)

	for i := 0; i < maxSleepTime; i++ {
		db, err := sqlx.Connect("mysql", dbConfigStr)
		if err == nil {
			DB = db
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		panic(err.Error())
	}

	return db
}
