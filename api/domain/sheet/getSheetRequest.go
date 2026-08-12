package sheet

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/domain/d7d"
	"github.com/aleciosouza/dm-helper/handler"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSheetHandler(ctx *gin.Context) {
	userID := ctx.GetUint(config.ContextUserID)
	sheetID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handler.SendError(ctx, http.StatusBadRequest, handler.ErrParamIsRequired("sheetID", "uint").Error())
		return
	}

	sheet := schemas.Sheet{}
	err = db.Where("id = ? AND user_id = ?", sheetID, userID).Preload("Skills").Preload("Attacks").First(&sheet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.SendError(ctx, http.StatusNotFound, "sheet not found")
			return
		}

		logger.Errorf("error loading sheet %d: %v", sheetID, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to get sheet: [SH-01]")
		return
	}

	handler.SendSuccess(ctx, "get_sheet", sheet.ToResponse())
}

func GetSheetsByUserHandler(ctx *gin.Context) {
	userID := ctx.GetUint(config.ContextUserID)

	sheets := []schemas.Sheet{}
	err := db.Where("user_id = ?", userID).Preload("Skills").Preload("Attacks").Find(&sheets).Error

	if err != nil {
		logger.Errorf("error loading sheets for user %d: %v", userID, err)
		handler.SendError(ctx, http.StatusInternalServerError, "failed to get sheets: [SH-02]")
		return
	}

	responses := make([]schemas.SheetResponse, len(sheets))
	for i, sheet := range sheets {
		responses[i] = sheet.ToResponse()
	}

	handler.SendSuccess(ctx, "get_sheets", responses)
}

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
