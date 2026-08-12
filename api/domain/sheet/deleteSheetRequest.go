package sheet

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteSheetHandler(ctx *gin.Context) {
	userId := ctx.GetUint(config.ContextUserID)
	sheetID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("sheetID", "uint").Error())
		return
	}

	sheet := schemas.Sheet{}
	err = db.Where("id = ? AND user_id = ?", sheetID, userId).First(&sheet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(ctx, http.StatusNotFound, "sheet not found")
			return
		}

		logger.Errorf("error loading sheet %d: %v", sheetID, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to get sheet: [SH-01]")
		return
	}

	if err = db.Delete(&sheet).Error; err != nil {
		logger.Errorf("error deleting sheet %d for user %d: %v", sheetID, userId, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to delete sheet: [SH-02]")
		return
	}

	handler.SendSuccess(ctx, "delete_sheet", gin.H{
		"sheet_id": sheet.ID,
	})
}
