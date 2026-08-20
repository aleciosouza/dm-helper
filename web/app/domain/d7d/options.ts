import { ESkill, type TDice } from "./types";

export interface IClassOption {
    value: string;
    hitDice: TDice;
}

export const CLASSES: readonly IClassOption[] = [
    { value: "Barbarian", hitDice: "d12" },
    { value: "Bard", hitDice: "d8" },
    { value: "Cleric", hitDice: "d8" },
    { value: "Druid", hitDice: "d8" },
    { value: "Fighter", hitDice: "d10" },
    { value: "Monk", hitDice: "d8" },
    { value: "Paladin", hitDice: "d10" },
    { value: "Ranger", hitDice: "d10" },
    { value: "Rogue", hitDice: "d8" },
    { value: "Sorcerer", hitDice: "d6" },
    { value: "Warlock", hitDice: "d8" },
    { value: "Wizard", hitDice: "d6" },
];

export interface IRaceOption {
    value: string;
    speed: string;
}

export const RACES: readonly IRaceOption[] = [
    { value: "Dragonborn", speed: "9m" },
    { value: "Dwarf", speed: "7.5m" },
    { value: "Elf", speed: "9 m" },
    { value: "Gnome", speed: "7.5m" },
    { value: "Half-Elf", speed: "9m" },
    { value: "Half-Orc", speed: "9m" },
    { value: "Halfling", speed: "7.5m" },
    { value: "Human", speed: "9m" },
    { value: "Tiefling", speed: "9m" },
];

export const ALIGNMENTS: readonly string[] = [
    "Lawful Good",
    "Neutral Good",
    "Chaotic Good",
    "Lawful Neutral",
    "Neutral",
    "Chaotic Neutral",
    "Lawful Evil",
    "Neutral Evil",
    "Chaotic Evil",
];

export interface ISkillOption {
    value: ESkill;
    label: string;
}

export const SKILLS: readonly ISkillOption[] = [
    { value: ESkill.Acrobatics, label: "Acrobatics" },
    { value: ESkill.AnimalHandling, label: "Animal Handling" },
    { value: ESkill.Arcana, label: "Arcana" },
    { value: ESkill.Athletics, label: "Athletics" },
    { value: ESkill.Deception, label: "Deception" },
    { value: ESkill.History, label: "History" },
    { value: ESkill.Insight, label: "Insight" },
    { value: ESkill.Intimidation, label: "Intimidation" },
    { value: ESkill.Investigation, label: "Investigation" },
    { value: ESkill.Medicine, label: "Medicine" },
    { value: ESkill.Nature, label: "Nature" },
    { value: ESkill.Perception, label: "Perception" },
    { value: ESkill.Performance, label: "Performance" },
    { value: ESkill.Persuasion, label: "Persuasion" },
    { value: ESkill.Religion, label: "Religion" },
    { value: ESkill.SleightOfHand, label: "Sleight of Hand" },
    { value: ESkill.Stealth, label: "Stealth" },
    { value: ESkill.Survival, label: "Survival" },
];

export const DICE: readonly TDice[] = ["d4", "d6", "d8", "d10", "d12", "d20", "d100"];

export interface IAbilityOption {
    key: "str" | "dex" | "con" | "int" | "wis" | "cha";
    save: "save_str" | "save_dex" | "save_con" | "save_int" | "save_wis" | "save_cha";
    label: string;
}

export const ABILITIES: readonly IAbilityOption[] = [
    { key: "str", save: "save_str", label: "Strength" },
    { key: "dex", save: "save_dex", label: "Dexterity" },
    { key: "con", save: "save_con", label: "Constitution" },
    { key: "int", save: "save_int", label: "Intelligence" },
    { key: "wis", save: "save_wis", label: "Wisdom" },
    { key: "cha", save: "save_cha", label: "Charisma" },
];
