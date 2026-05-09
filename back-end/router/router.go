package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	InitRoutes(r)

	r.Run(":8080")
}
