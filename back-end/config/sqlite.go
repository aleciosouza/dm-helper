package config

import (
	"os"

	"github.com/aleciosouza/dm-helper/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_PATH = "./db/main.db"

func checkDatabaseFile() error {
	_, err := os.Stat(DB_PATH)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating new database")

		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Failed to create database directory: %v", err)
			return err
		}

		file, err := os.Create(DB_PATH)

		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return err
		}

		file.Close()
	}

	return nil
}

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Check if database file exists
	err := checkDatabaseFile()
	if err != nil {
		return nil, err
	}

	// Create database and connect
	db, err := gorm.Open(
		sqlite.Open(DB_PATH),
		&gorm.Config{},
	)

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Automigrate schemas
	err = db.AutoMigrate(&schemas.Sheet{})
	if err != nil {
		logger.Errorf("SQLite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
