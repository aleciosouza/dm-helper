package handler

import (
	"net/http"

	"github.com/aleciosouza/dm-helper/config"
	"github.com/aleciosouza/dm-helper/d7d"
	"github.com/aleciosouza/dm-helper/schemas"
	"github.com/gin-gonic/gin"
)

type CreateSheetRequest struct {
	d7d.Sheet
}

type SheetSkill struct {
	Skill      d7d.Skill `json:"skill"`
	Proficient bool      `json:"proficient"`
}

type SheetAttack struct {
	Name             string   `json:"name"`
	AttackDice       d7d.Dice `json:"attack_dice"`
	AttackDiceAmount uint8    `json:"attack_dice_amount"`
	DamageType       string   `json:"damage_type"`
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

	logger := config.GetLogger("createSheet")
	userID := ctx.GetUint(config.ContextUserID)
	logger.Infof("Creating sheet for user: %v", userID)

	sheet := schemas.Sheet{
		UserID:             userID,
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

	if err := db.Create(&sheet).Error; err != nil {
		logger.Errorf("Error creating sheet: %v", err)
		sendError(ctx, http.StatusInternalServerError, "failed to create sheet: [SH-01]")
		return
	}

	sendSuccess(ctx, "create_sheet", gin.H{
		"sheet_id": sheet.ID,
	})
}
