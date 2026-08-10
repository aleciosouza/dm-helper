package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/d7d"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func sheetFromRequest(userId uint, req d7d.Sheet) schemas.Sheet {
	skills := make([]schemas.SheetSkill, len(req.Skills))
	for i, skill := range req.Skills {
		skills[i] = schemas.SheetSkill{
			Skill: skill,
		}
	}

	attacks := make([]schemas.SheetAttack, len(req.Attacks))
	for i, attack := range req.Attacks {
		attacks[i] = schemas.SheetAttack{
			Name:             attack.Name,
			AttackDice:       attack.AttackDice,
			AttackDiceAmount: attack.AttackDiceAmount,
			DamageType:       attack.DamageType,
		}
	}

	sheet := schemas.Sheet{
		UserID:             userId,
		Name:               req.Name,
		Class:              req.Class,
		Race:               req.Race,
		Background:         req.Background,
		Alignment:          req.Alignment,
		PersonalityTraits:  req.PersonalityTraits,
		Ideals:             req.Ideals,
		Bonds:              req.Bonds,
		Flaws:              req.Flaws,
		XP:                 req.XP,
		Level:              req.Level,
		STR:                req.STR,
		DEX:                req.DEX,
		CON:                req.CON,
		INT:                req.INT,
		WIS:                req.WIS,
		CHA:                req.CHA,
		SaveSTR:            req.SaveSTR,
		SaveDEX:            req.SaveDEX,
		SaveCON:            req.SaveCON,
		SaveINT:            req.SaveINT,
		SaveWIS:            req.SaveWIS,
		SaveCHA:            req.SaveCHA,
		ArmorClass:         req.ArmorClass,
		Speed:              req.Speed,
		Hp:                 req.Hp,
		HpMax:              req.HpMax,
		HpTemp:             req.HpTemp,
		HitDice:            req.HitDice,
		HitDiceCurrent:     req.HitDiceCurrent,
		HitDiceTotal:       req.HitDiceTotal,
		DeathSavesSuccess:  req.DeathSavesSuccess,
		DeathSavesFailure:  req.DeathSavesFailure,
		Languages:          req.Languages,
		OtherProficiencies: req.OtherProficiencies,
		Age:                req.Age,
		Height:             req.Height,
		Weight:             req.Weight,
		Eyes:               req.Eyes,
		Skin:               req.Skin,
		Hair:               req.Hair,
		Appearance:         req.Appearance,
		Backstory:          req.Backstory,
		Skills:             skills,
		Attacks:            attacks,
	}
	return sheet
}

type CreateSheetRequest struct {
	d7d.Sheet
}

func (s *CreateSheetRequest) Validate() error {
	if s.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if s.Class == "" {
		return errParamIsRequired("class", "string")
	}

	if s.Race == "" {
		return errParamIsRequired("race", "string")
	}

	if s.Level < 1 {
		return errParamIsRequired("level", "uint8")
	}

	if s.Speed == "" {
		return errParamIsRequired("Speed", "string")
	}

	if s.HitDice == "" {
		return errParamIsRequired("HitDice", "string")
	}

	return nil
}

func CreateSheetHandler(ctx *gin.Context) {
	req := CreateSheetRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userID := ctx.GetUint(config.ContextUserID)

	sheet := sheetFromRequest(userID, req.Sheet)

	if err := db.Create(&sheet).Error; err != nil {
		logger.Errorf("Error creating sheet: %v", err)
		sendError(ctx, http.StatusInternalServerError, "failed to create sheet: [SH-01]")
		return
	}

	sendSuccess(ctx, "create_sheet", gin.H{
		"sheet_id": sheet.ID,
	})
}

type UpdateSheetRequest struct {
	d7d.Sheet
}

func (s *UpdateSheetRequest) Validate() error {
	return nil
}

func UpdateSheetHandler(ctx *gin.Context) {
	sheetID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("sheetID", "uint").Error())
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	userID := ctx.GetUint(config.ContextUserID)

	current := schemas.Sheet{}
	if err := db.Preload("Skills").Preload("Attacks").
		Where("id = ? AND user_id = ?", sheetID, userID).
		First(&current).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "sheet not found")
			return
		}

		logger.Errorf("error loading sheet %d: %v", sheetID, err)
		sendError(ctx, http.StatusInternalServerError, "failed to update sheet: [SH-02]")
		return
	}

	req := UpdateSheetRequest{Sheet: current.ToResponse().Sheet}
	if err := json.Unmarshal(body, &req); err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sheet := sheetFromRequest(userID, req.Sheet)
	sheet.ID = current.ID
	sheet.CreatedAt = current.CreatedAt

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Skills", "Attacks").Save(&sheet).Error; err != nil {
			return err
		}

		if len(req.Skills) > 0 && len(sheet.Skills) > 0 {
			if err := tx.Where("sheet_id = ?", sheet.ID).Delete(&schemas.SheetSkill{}).Error; err != nil {
				return err
			}

			for i := range sheet.Skills {
				sheet.Skills[i].SheetID = sheet.ID
			}

			if err := tx.Create(&sheet.Skills).Error; err != nil {
				return err
			}
		}

		if len(req.Attacks) > 0 && len(sheet.Attacks) > 0 {
			if err := tx.Where("sheet_id = ?", sheet.ID).Delete(&schemas.SheetAttack{}).Error; err != nil {
				return err
			}

			for i := range sheet.Attacks {
				sheet.Attacks[i].SheetID = sheet.ID
			}
			if err := tx.Create(&sheet.Attacks).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logger.Errorf("error updating sheet %d for user %d: %v", sheetID, userID, err)
		sendError(ctx, http.StatusInternalServerError, "failed to update sheet: [SH-02]")
		return
	}

	sendSuccess(ctx, "update_sheet", sheet.ToResponse())
}
