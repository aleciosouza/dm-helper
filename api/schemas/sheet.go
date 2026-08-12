package schemas

import (
	"time"

	"github.com/aleciosouza/dm-helper/domain/d7d"
)

type SheetSkill struct {
	SheetID uint      `gorm:"primaryKey"`
	Skill   d7d.Skill `gorm:"primaryKey"`
}

type SheetAttack struct {
	SheetID          uint     `gorm:"primaryKey"`
	Name             string   `gorm:"primaryKey"`
	AttackDice       d7d.Dice `gorm:"not null"`
	AttackDiceAmount uint8    `gorm:"not null"`
	DamageType       string   `gorm:"not null"`
}

type Sheet struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID            uint   `gorm:"not null;index"`
	Name              string `gorm:"not null"`
	Class             string `gorm:"not null"`
	Race              string `gorm:"not null"`
	Background        string
	Alignment         string
	PersonalityTraits string
	Ideals            string
	Bonds             string
	Flaws             string

	XP    uint32 `gorm:"default:0"`
	Level uint8  `gorm:"default:1"`

	STR uint8 `gorm:"not null"`
	DEX uint8 `gorm:"not null"`
	CON uint8 `gorm:"not null"`
	INT uint8 `gorm:"not null"`
	WIS uint8 `gorm:"not null"`
	CHA uint8 `gorm:"not null"`

	SaveSTR bool `gorm:"default:false"`
	SaveDEX bool `gorm:"default:false"`
	SaveCON bool `gorm:"default:false"`
	SaveINT bool `gorm:"default:false"`
	SaveWIS bool `gorm:"default:false"`
	SaveCHA bool `gorm:"default:false"`

	ArmorClass uint8  `gorm:"not null"`
	Speed      string `gorm:"not null"`

	Hp     uint8 `gorm:"not null"`
	HpMax  uint8 `gorm:"not null"`
	HpTemp uint8 `gorm:"default:0"`

	HitDice        string `gorm:"not null"`
	HitDiceCurrent uint8  `gorm:"not null"`
	HitDiceTotal   uint8  `gorm:"not null"`

	DeathSavesSuccess uint8 `gorm:"default:0"`
	DeathSavesFailure uint8 `gorm:"default:0"`

	Skills  []SheetSkill  `gorm:"foreignKey:SheetID;constraint:OnDelete:CASCADE"`
	Attacks []SheetAttack `gorm:"foreignKey:SheetID;constraint:OnDelete:CASCADE"`

	Languages          string
	OtherProficiencies string

	Age        uint8
	Height     string
	Weight     string
	Eyes       string
	Skin       string
	Hair       string
	Appearance string
	Backstory  string
}

type SheetResponse struct {
	d7d.Sheet
	ID uint `json:"id"`
}

func (s *Sheet) ToResponse() SheetResponse {
	attacks := make([]d7d.SheetAttack, len(s.Attacks))

	for i, attack := range s.Attacks {
		attacks[i] = d7d.SheetAttack{
			Name:             attack.Name,
			AttackDice:       attack.AttackDice,
			AttackDiceAmount: attack.AttackDiceAmount,
			DamageType:       attack.DamageType,
		}
	}

	skills := make([]d7d.Skill, len(s.Skills))

	for i, skill := range s.Skills {
		skills[i] = skill.Skill
	}

	return SheetResponse{
		ID: s.ID,
		Sheet: d7d.Sheet{
			Name:              s.Name,
			Class:             s.Class,
			Race:              s.Race,
			Background:        s.Background,
			Alignment:         s.Alignment,
			PersonalityTraits: s.PersonalityTraits,
			Ideals:            s.Ideals,
			Bonds:             s.Bonds,
			Flaws:             s.Flaws,

			XP:    s.XP,
			Level: s.Level,

			STR: s.STR,
			DEX: s.DEX,
			CON: s.CON,
			INT: s.INT,
			WIS: s.WIS,
			CHA: s.CHA,

			SaveSTR: s.SaveSTR,
			SaveDEX: s.SaveDEX,
			SaveCON: s.SaveCON,
			SaveINT: s.SaveINT,
			SaveWIS: s.SaveWIS,
			SaveCHA: s.SaveCHA,

			Skills: skills,

			ArmorClass: s.ArmorClass,
			Speed:      s.Speed,

			Hp:     s.Hp,
			HpMax:  s.HpMax,
			HpTemp: s.HpTemp,

			HitDice:        s.HitDice,
			HitDiceCurrent: s.HitDiceCurrent,
			HitDiceTotal:   s.HitDiceTotal,

			DeathSavesSuccess: s.DeathSavesSuccess,
			DeathSavesFailure: s.DeathSavesFailure,

			Attacks: attacks,

			Languages:          s.Languages,
			OtherProficiencies: s.OtherProficiencies,

			Age:        s.Age,
			Height:     s.Height,
			Weight:     s.Weight,
			Eyes:       s.Eyes,
			Skin:       s.Skin,
			Hair:       s.Hair,
			Appearance: s.Appearance,
			Backstory:  s.Backstory,
		},
	}
}
