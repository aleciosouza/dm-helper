package schemas

import "gorm.io/gorm"

// D&D 5e character sheet schema
type Sheet struct {
	gorm.Model
	Name     string `json:"name"`
	PlayerId uint   `json:"playerId"`
	Class    string `json:"class"`
	Level    int    `json:"level"`
	DEX      int    `json:"dex"`
	CON      int    `json:"con"`
	INT      int    `json:"int"`
	WIS      int    `json:"wis"`
	CHA      int    `json:"cha"`
	STR      int    `json:"str"`
}
