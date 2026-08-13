package auth

import (
	"net/http"
	"strings"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/gin-gonic/gin"
)

const BEARER_PREFIX = "bearer "

func AuthHandler(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		handler.SendError(ctx, http.StatusUnauthorized, "missing authorization header")
		return
	}

	if len(header) < len(BEARER_PREFIX) || !strings.EqualFold(header[:len(BEARER_PREFIX)], BEARER_PREFIX) {
		handler.SendError(ctx, http.StatusNotAcceptable, "invalid authorization header")
		return
	}

	token := strings.TrimSpace(header[len(BEARER_PREFIX):])

	claims, err := config.ValidateJWT(token)

	if err != nil {
		handler.SendError(ctx, http.StatusUnauthorized, "invalid token")
		return
	}

	ctx.Set(config.ContextUserID, claims.UserID)
	ctx.Status(http.StatusOK)
}
