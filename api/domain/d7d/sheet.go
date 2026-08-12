package d7d

type Sheet struct {
	Name              string `json:"name"`
	Class             string `json:"class"`
	Race              string `json:"race"`
	Background        string `json:"background"`
	Alignment         string `json:"alignment"`
	PersonalityTraits string `json:"personality_traits"`
	Ideals            string `json:"ideals"`
	Bonds             string `json:"bonds"`
	Flaws             string `json:"flaws"`

	XP    uint32 `json:"xp"`
	Level uint8  `json:"level"`

	STR uint8 `json:"str"`
	DEX uint8 `json:"dex"`
	CON uint8 `json:"con"`
	INT uint8 `json:"int"`
	WIS uint8 `json:"wis"`
	CHA uint8 `json:"cha"`

	SaveSTR bool `json:"save_str"`
	SaveDEX bool `json:"save_dex"`
	SaveCON bool `json:"save_con"`
	SaveINT bool `json:"save_int"`
	SaveWIS bool `json:"save_wis"`
	SaveCHA bool `json:"save_cha"`

	Skills []Skill `json:"skills"`

	ArmorClass uint8  `json:"armor_class"`
	Speed      string `json:"speed"`

	Hp     uint8 `json:"hp"`
	HpMax  uint8 `json:"hp_max"`
	HpTemp uint8 `json:"hp_temp"`

	HitDice        string `json:"hit_dice"`
	HitDiceCurrent uint8  `json:"hit_dice_current"`
	HitDiceTotal   uint8  `json:"hit_dice_total"`

	DeathSavesSuccess uint8 `json:"death_saves_success"`
	DeathSavesFailure uint8 `json:"death_saves_failure"`

	Attacks []SheetAttack `json:"attacks"`

	Languages          string `json:"languages"`
	OtherProficiencies string `json:"other_proficiencies"`

	Age        uint8  `json:"age"`
	Height     string `json:"height"`
	Weight     string `json:"weight"`
	Eyes       string `json:"eyes"`
	Skin       string `json:"skin"`
	Hair       string `json:"hair"`
	Appearance string `json:"appearance"`
	Backstory  string `json:"backstory"`
}

type SheetAttack struct {
	Name             string `json:"name"`
	AttackDice       Dice   `json:"attack_dice"`
	AttackDiceAmount uint8  `json:"attack_dice_amount"`
	DamageType       string `json:"damage_type"`
}
