package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSheetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE sheet",
	})
}
