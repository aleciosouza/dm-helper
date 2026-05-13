package handler

import (
	"net/http"

	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
)

func CreateSheetHandler(ctx *gin.Context) {
	request := CreateSheetRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sheet := schemas.Sheet{
		Name:     request.Name,
		PlayerId: request.PlayerId,
		Class:    request.Class,
		Level:    request.Level,
		DEX:      request.DEX,
		CON:      request.CON,
		INT:      request.INT,
		WIS:      request.WIS,
		CHA:      request.CHA,
		STR:      request.STR,
	}

	if err := db.Create(&sheet).Error; err != nil {
		logger.Errorf("Error creating sheet: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Failed to create sheet")
		return
	}

	sendSuccess(ctx, "create-sheet", sheet)
}
