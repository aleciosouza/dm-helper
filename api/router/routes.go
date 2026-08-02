package router

import (
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	handler.InitHandler()

	v1 := router.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/register", handler.RegisterHandler)
		auth.POST("/login", handler.LoginHandler)
	}
}
