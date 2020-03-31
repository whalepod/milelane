package handler

import (
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var now = time.Now()

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("sqlmock", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}
