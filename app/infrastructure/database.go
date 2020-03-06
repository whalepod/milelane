package infrastructure

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// To enable gorm to connect MySQL.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB exposes database connection.
var DB *gorm.DB

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

func connectDB() *gorm.DB {
	// Initalize db connection.
	m := mySQLConfig{
		Username: os.Getenv("MILELANE_DATABASE_USERNAME"),
		Password: os.Getenv("MILELANE_DATABASE_PASSWORD"),
		Host:     os.Getenv("MILELANE_DATABASE_HOST"),
		Database: os.Getenv("MILELANE_DATABASE"),
	}

	dbConfigStr := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":3306)/" + m.Database + "?parseTime=true"
	db, err := gorm.Open("mysql", dbConfigStr)

	for i := 0; i < maxSleepTime; i++ {
		db, err := gorm.Open("mysql", dbConfigStr)
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
