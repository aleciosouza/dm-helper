import type { DefaultValues } from "react-hook-form";

import { CLASSES, RACES } from "../options";
import type { ISheet } from "./sheet";

export type TSheetForm = Omit<ISheet, "id">;

export function defaultSheetForm(): DefaultValues<TSheetForm> {
    return {
        name: "",
        class: "",
        race: "",
        background: "",
        alignment: "",
        personality_traits: "",
        ideals: "",
        bonds: "",
        flaws: "",

        xp: 0,
        level: 1,

        str: undefined,
        dex: undefined,
        con: undefined,
        int: undefined,
        wis: undefined,
        cha: undefined,

        save_str: false,
        save_dex: false,
        save_con: false,
        save_int: false,
        save_wis: false,
        save_cha: false,

        skills: [],

        armor_class: 0,
        speed: "",

        hp: 0,
        hp_max: 0,
        hp_temp: 0,

        hit_dice: "d8",
        hit_dice_current: 1,
        hit_dice_total: 1,

        death_saves_success: 0,
        death_saves_failure: 0,

        attacks: [],

        languages: "",
        other_proficiencies: "",

        age: undefined,
        height: "",
        weight: "",
        eyes: "",
        skin: "",
        hair: "",
        appearance: "",
        backstory: "",
    };
}

export function toSheetForm(sheet: ISheet): TSheetForm {
    const { id: _id, ...values } = sheet;

    return values;
}

export function toSheetPayload(values: TSheetForm): TSheetForm {
    const race = RACES.find((option) => option.value === values.race);
    const characterClass = CLASSES.find((option) => option.value === values.class);

    return {
        ...values,
        age: Number.isNaN(values.age) ? 0 : values.age,
        speed: race?.speed ?? "9m",
        hit_dice: characterClass?.hitDice ?? "d8",
        hit_dice_current: values.level,
        hit_dice_total: values.level,
    };
}
