package sheet

import (
	"github.com/aleciosouza/dm-helper/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitModule() {
	logger = config.GetLogger("sheet")
	db = config.GetDB()
}
