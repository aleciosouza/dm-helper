package handler

import "fmt"

func errParamIsRequired(name string, paramType string) error {
	return fmt.Errorf("param: %s (%s) is required", name, paramType)
}

// Create Sheet
type CreateSheetRequest struct {
	Name     string `json:"name"`
	PlayerId uint   `json:"playerId"`
	Class    string `json:"class"`
	Level    uint8  `json:"level"`
	DEX      uint8  `json:"dex"`
	CON      uint8  `json:"con"`
	INT      uint8  `json:"int"`
	WIS      uint8  `json:"wis"`
	CHA      uint8  `json:"cha"`
	STR      uint8  `json:"str"`
}

func (r *CreateSheetRequest) Validate() error {
	if r.Name == "" &&
		r.PlayerId == 0 &&
		r.Class == "" &&
		r.Level == 0 &&
		r.DEX == 0 &&
		r.CON == 0 &&
		r.INT == 0 &&
		r.WIS == 0 &&
		r.CHA == 0 &&
		r.STR == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.PlayerId == 0 {
		return errParamIsRequired("playerId", "uint")
	}

	if r.Class == "" {
		return errParamIsRequired("class", "string")
	}

	if r.Level == 0 {
		return errParamIsRequired("level", "uint8")
	}

	if r.DEX == 0 {
		return errParamIsRequired("dex", "uint8")
	}

	if r.CON == 0 {
		return errParamIsRequired("con", "uint8")
	}

	if r.INT == 0 {
		return errParamIsRequired("int", "uint8")
	}

	if r.WIS == 0 {
		return errParamIsRequired("wis", "uint8")
	}

	if r.CHA == 0 {
		return errParamIsRequired("cha", "uint8")
	}

	if r.STR == 0 {
		return errParamIsRequired("str", "uint8")
	}

	return nil
}
