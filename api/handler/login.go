package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

func LoginHandler(ctx *gin.Context) {
	req := LoginRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusUnauthorized, "invalid credentials")
			return
		}

		logger.Errorf("Error fetching user: %v", err)
		sendError(ctx, http.StatusInternalServerError, "failed to login: [LG-01]")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := config.GenerateJWT(user.ID)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to login: [LG-02]")
		return
	}

	sendSuccess(ctx, "login", gin.H{
		"token": token,
		"user":  user.ToResponse(),
	})
}
