package main

import (
	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/router"
)

func main() {
	config.LoadEnv()

	logger := config.GetLogger("main")

	if err := config.InitDB(); err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.InitRouter()
}
