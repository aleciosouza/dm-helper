package main

import "github.com/aleciosouza/dm-helper/config"

func main() {
	config.LoadEnv()

	logger = config.GetLogger("main")
}
