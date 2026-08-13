package router

import (
	"github.com/aleciosouza/dm-helper/domain/auth"
	"github.com/aleciosouza/dm-helper/domain/sheet"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	handler.InitHandler()

	v1 := router.Group("api/v1")

	auth.InitModule()
	authGroup := v1.Group("auth")
	{
		authGroup.HEAD("", auth.AuthHandler)
		authGroup.POST("/register", auth.RegisterHandler)
		authGroup.POST("/login", auth.LoginHandler)
	}

	sheet.InitModule()
	sheetGroup := v1.Group("sheet", AuthMiddleware())
	{
		sheetGroup.GET("", sheet.GetSheetsByUserHandler)
		sheetGroup.GET("/:id", sheet.GetSheetHandler)
		sheetGroup.POST("", sheet.CreateSheetHandler)
		sheetGroup.PATCH("/:id", sheet.UpdateSheetHandler)
		sheetGroup.DELETE("/:id", sheet.DeleteSheetHandler)
	}
}
