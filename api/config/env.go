package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	logger := NewLogger("env")

	if err := godotenv.Load(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Warningf("Could not load .env: %v", err)
		}

		logger.Warningf("No .env file found")
		return
	}

	logger.Info("Loaded .env")
}
