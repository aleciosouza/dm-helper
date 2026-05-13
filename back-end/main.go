package main

import (
	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize the database connection
	if err := config.InitDB(); err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.InitRouter()
}
