package sheet

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/domain/d7d"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateSheetRequest struct {
	d7d.Sheet
}

func (s *UpdateSheetRequest) Validate() error {
	return nil
}

func UpdateSheetHandler(ctx *gin.Context) {
	sheetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("sheetID", "uint").Error())
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	userID := ctx.GetUint(config.ContextUserID)

	current := schemas.Sheet{}
	if err := db.Preload("Skills").Preload("Attacks").
		Where("id = ? AND user_id = ?", sheetID, userID).
		First(&current).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(ctx, http.StatusNotFound, "sheet not found")
			return
		}

		logger.Errorf("error loading sheet %d: %v", sheetID, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to update sheet: [SH-01]")
		return
	}

	req := UpdateSheetRequest{Sheet: current.ToResponse().Sheet}
	if err := json.Unmarshal(body, &req); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sheet := sheetFromRequest(userID, req.Sheet)
	sheet.ID = current.ID
	sheet.CreatedAt = current.CreatedAt

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Skills", "Attacks").Save(&sheet).Error; err != nil {
			return err
		}

		if err := tx.Where("sheet_id = ?", sheet.ID).Delete(&schemas.SheetSkill{}).Error; err != nil {
			return err
		}

		if len(sheet.Skills) > 0 {
			for i := range sheet.Skills {
				sheet.Skills[i].SheetID = sheet.ID
			}

			if err := tx.Create(&sheet.Skills).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("sheet_id = ?", sheet.ID).Delete(&schemas.SheetAttack{}).Error; err != nil {
			return err
		}

		if len(sheet.Attacks) > 0 {
			for i := range sheet.Attacks {
				sheet.Attacks[i].SheetID = sheet.ID
			}

			if err := tx.Create(&sheet.Attacks).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logger.Errorf("error updating sheet %d for user %d: %v", sheetID, userID, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to update sheet: [SH-02]")
		return
	}

	handler.SendSuccess(ctx, "update_sheet", sheet.ToResponse())
}
