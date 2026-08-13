package router

import (
	"net/http"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/domain/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, bearerError := auth.GetBearer(ctx)

		if bearerError != nil {
			ctx.AbortWithStatusJSON(bearerError.Status, gin.H{"error": bearerError.Error()})
			return
		}

		claims, err := config.ValidateJWT(token)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx.Set(config.ContextUserID, claims.UserID)
		ctx.Next()
	}
}
