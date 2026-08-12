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

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	r.Name = strings.TrimSpace(r.Name)

	if r.Email == "" {
		return handler.ErrParamIsRequired("email", "string")
	}

	if !strings.Contains(r.Email, "@") {
		return errors.New("invalid email")
	}

	if len(r.Password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	if r.Name == "" {
		return handler.ErrParamIsRequired("name", "string")
	}

	return nil
}

func RegisterHandler(ctx *gin.Context) {
	req := RegisterRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	existing := schemas.User{}
	err := db.Where("email = ?", req.Email).First(&existing).Error
	if err == nil {
		handler.SendError(ctx, http.StatusConflict, "email already registered")
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("Error to check if email is available: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to register: [RG-01]")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("Error hashing password: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to register: [RG-02]")
		return
	}

	user := schemas.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("Error creating user: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to register: [RG-03]")
		return
	}

	token, err := config.GenerateJWT(user.ID)

	if err != nil {
		logger.Errorf("Error generating token: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to register: [RG-04]")
		return
	}

	handler.SendSuccess(ctx, "auth:register", gin.H{
		"token": token,
		"user":  user.ToResponse(),
	})
}
