package infrastructure

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

const maxSleepTime = 30

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

func init() {
	// Start db connection.
	var config Config

	var configFilePath string
	basePath := os.Getenv("MILELANE_PATH")

	env := os.Getenv("MILELANE_ENV")
	switch env {
	case "development":
		configFilePath = basePath + "config/development.toml"
	case "production":
		configFilePath = basePath + "config/production.toml"
	case "test":
		configFilePath = basePath + "config/test.toml"
	default:
		configFilePath = basePath + "config/development.toml"
	}

	_, err := toml.DecodeFile(configFilePath, &config)
	if err != nil {
		panic(err.Error())
	}

	dbConfig := config.MySQL.Username + ":" + config.MySQL.Password + "@tcp(" + config.MySQL.Host + ":3306)/" + config.MySQL.Database + "?parseTime=true"
	db, err := gorm.Open("mysql", dbConfig)

	for i := 0; i < maxSleepTime; i++ {
		db, err := gorm.Open("mysql", dbConfig)
		if err == nil {
			DB = db
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
