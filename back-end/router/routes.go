package router

import (
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/sheet", handler.GetSheetHandler)
		v1.POST("/sheet", handler.CreateSheetHandler)
		v1.DELETE("/sheet", handler.DeleteSheetHandler)
		v1.PUT("/sheet", handler.UpdateSheetHandler)
		v1.GET("/sheets", handler.ListSheetsHandler)
	}
}
