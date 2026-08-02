package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	InitRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
