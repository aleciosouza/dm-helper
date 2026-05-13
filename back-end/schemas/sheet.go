package schemas

import (
	"time"

	"gorm.io/gorm"
)

// D&D 5e character sheet schema
type Sheet struct {
	gorm.Model
	Name     string
	PlayerId uint
	Class    string
	Level    uint8
	DEX      uint8
	CON      uint8
	INT      uint8
	WIS      uint8
	CHA      uint8
	STR      uint8
}

type SheetResponse struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	PlayerId  uint      `json:"playerId"`
	Class     string    `json:"class"`
	Level     uint8     `json:"level"`
	DEX       uint8     `json:"dex"`
	CON       uint8     `json:"con"`
	INT       uint8     `json:"int"`
	WIS       uint8     `json:"wis"`
	CHA       uint8     `json:"cha"`
	STR       uint8     `json:"str"`
}
