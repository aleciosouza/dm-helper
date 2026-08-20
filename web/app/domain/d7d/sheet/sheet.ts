import type { ESkill, TDice } from "../types";

export interface ISheetAttack {
    name: string;
    attack_dice: TDice;
    attack_dice_amount: number;
    damage_type: string;
}

export interface ISheet {
    id: number
    name: string
    class: string
    race: string
    background: string
    alignment: string
    personality_traits: string
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

    skills: ESkill[]

    armor_class: number
    speed: string

    hp: number
    hp_max: number
    hp_temp: number

    hit_dice: TDice
    hit_dice_current: number
    hit_dice_total: number

    death_saves_success: number
    death_saves_failure: number

    attacks: ISheetAttack[]

    languages: string
    other_proficiencies: string

    age: number
    height: string
    weight: string
    eyes: string
    skin: string
    hair: string
    appearance: string
    backstory: string
}
