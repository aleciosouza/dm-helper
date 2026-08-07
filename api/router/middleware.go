package router

import (
	"net/http"
	"strings"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/gin-gonic/gin"
)

const BEARER_PREFIX = "bearer "

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		if len(header) < len(BEARER_PREFIX) || !strings.EqualFold(header[:len(BEARER_PREFIX)], BEARER_PREFIX) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		token := strings.TrimSpace(header[len(BEARER_PREFIX):])

		claims, err := config.ValidateJWT(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set(config.ContextUserID, claims.UserID)
		ctx.Next()
	}
}
