package auth

import (
	"net/http"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context) {
	token, bearerError := GetBearer(ctx)

	if bearerError != nil {
		handler.SendError(ctx, bearerError.Status, bearerError.Error())
		return
	}

	claims, err := config.ValidateJWT(token)

	if err != nil {
		handler.SendError(ctx, http.StatusUnauthorized, "invalid token")
		return
	}

	ctx.Set(config.ContextUserID, claims.UserID)
	ctx.Status(http.StatusOK)
}
