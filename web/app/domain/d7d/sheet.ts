import type { TDice, TSkill } from "./types";

export interface ISheetAttack {
    name: string;
    attackDice: TDice;
    attackDiceAmount: number;
    damageType: string;
}

export interface ISheet {
    name: string
    class: string
    race: string
    background: string
    alignment: string
    personalityTraits: string
    ideals: string
    bonds: string
    flaws: string
    xp: number
    level: number

    str: number
    dex: number
    con: number
    int: number
    wis: number
    cha: number

    save_str: boolean
    save_dex: boolean
    save_con: boolean
    save_int: boolean
    save_wis: boolean
    save_cha: boolean

    skills: TSkill[]

    armor_class: number
    speed: string

    hp: number
    hp_max: number
    hp_temp: number

    hit_dice: TDice
    hit_dice_current: number
    hit_dice_total: number

    deathSavesSuccess: number
    deathSavesFailure: number

    attacks: ISheetAttack[]

    languages: string
    otherProficiencies: string

    age: number
    height: string
    weight: string
    eyes: string
    skin: string
    hair: string
    appearance: string
    backstory: string
}
