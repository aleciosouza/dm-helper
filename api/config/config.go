package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}

func GetDB() *gorm.DB {
	return db
}

func InitDB() error {
	var err error

	db, err = InitPostgres()

	if err != nil {
		return fmt.Errorf("error initializing Postgres: %v", err)
	}

	return nil
}
