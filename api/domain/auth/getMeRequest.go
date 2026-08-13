package auth

import (
	"errors"
	"net/http"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMeHandler(ctx *gin.Context) {
	userId := ctx.GetUint(config.ContextUserID)

	user := schemas.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(ctx, http.StatusNotFound, "user not found")
			return
		}

		logger.Errorf("Error fetching user: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed fetch user: [ME-01]")
		return
	}

	handler.SendSuccess(ctx, "login", user.ToResponse())
}
