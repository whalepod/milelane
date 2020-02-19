package repository

import (
	"database/sql/driver"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// AnyTime is value for sqlmock's `WithArgs` as timestamp,
// to avoid μ seconds difference of time.
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

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
