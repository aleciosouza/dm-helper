package auth

import (
	"net/http"
	"strings"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitModule() {
	logger = config.GetLogger("auth")
	db = config.GetDB()
}

const BEARER_PREFIX = "bearer "

type BearerError struct {
	Status  int
	message string
}

func (e *BearerError) Error() string {
	return e.message
}

func GetBearer(ctx *gin.Context) (string, *BearerError) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		return "", &BearerError{
			Status:  http.StatusUnauthorized,
			message: "missing authorization header",
		}
	}

	if len(bearer) < len(BEARER_PREFIX) || !strings.EqualFold(bearer[:len(BEARER_PREFIX)], BEARER_PREFIX) {
		return "", &BearerError{
			Status:  http.StatusNotAcceptable,
			message: "invalid authorization header",
		}
	}

	token := strings.TrimSpace(bearer[len(BEARER_PREFIX):])

	return token, nil
}
