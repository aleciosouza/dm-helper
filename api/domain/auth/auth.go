package auth

import (
	"github.com/aleciosouza/dm-helper/config"
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
