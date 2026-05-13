package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDB() error {
	var err error

	db, err = InitSQLite()

	if err != nil {
		return fmt.Errorf("Error initializing SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
