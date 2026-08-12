package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/handler"
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
		return handler.ErrParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return handler.ErrParamIsRequired("password", "string")
	}

	return nil
}

func LoginHandler(ctx *gin.Context) {
	req := LoginRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(ctx, http.StatusUnauthorized, "invalid credentials")
			return
		}

		logger.Errorf("Error fetching user: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to login: [LG-01]")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		handler.SendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := config.GenerateJWT(user.ID)

	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "failed to login: [LG-02]")
		return
	}

	handler.SendSuccess(ctx, "login", gin.H{
		"token": token,
		"user":  user.ToResponse(),
	})
}
