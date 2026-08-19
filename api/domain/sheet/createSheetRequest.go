package sheet

import (
	"net/http"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/domain/d7d"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

type CreateSheetRequest struct {
	d7d.Sheet
}

func (s *CreateSheetRequest) Validate() error {
	if s.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}

	if s.Class == "" {
		return handler.ErrParamIsRequired("class", "string")
	}

	if s.Race == "" {
		return handler.ErrParamIsRequired("race", "string")
	}

	if s.Level < 1 {
		return handler.ErrParamIsRequired("level", "uint8")
	}

	if s.Speed == "" {
		return handler.ErrParamIsRequired("Speed", "string")
	}

	if s.HitDice == "" {
		return handler.ErrParamIsRequired("HitDice", "string")
	}

	return nil
}

func CreateSheetHandler(ctx *gin.Context) {
	req := CreateSheetRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userID := ctx.GetUint(config.ContextUserID)

	sheet := sheetFromRequest(userID, req.Sheet)

	if err := db.Create(&sheet).Error; err != nil {
		logger.Errorf("Error creating sheet: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to create sheet: [SH-01]")
		return
	}

	handler.SendSuccess(ctx, "create_sheet", sheet.ToResponse())
}
