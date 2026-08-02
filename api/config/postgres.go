package config

import (
	"fmt"
	"os"

	"github.com/aleciosouza/dm-helper/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildDSN() (string, error) {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn, nil
	}

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if host == "" || user == "" || dbname == "" {
		return "", fmt.Errorf("no database configuration found: set DATABASE_URL or PGHOST/PGUSER/PGDATABASE")
	}

	if port == "" {
		port = "5432"
	}

	if sslmode == "" {
		sslmode = "disable"
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	), nil
}

func InitPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn, err := buildDSN()
	if err != nil {
		logger.Errorf("Failed to resolve database DSN: %v", err)
		return nil, err
	}

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := AutoMigrate(db); err != nil {
		logger.Errorf("Postgres automigration error: %v", err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&schemas.User{},
	)
}
