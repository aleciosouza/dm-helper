import { DICE } from "@/domain/d7d/options";
import type { TSheetForm } from "@/domain/d7d/sheet";
import { useFieldArray, type UseFormReturn } from "react-hook-form";

import Field from "./Field";

import styles from "./sheet.module.scss";

const StepAttacks: React.FC<{ form: UseFormReturn<TSheetForm> }> = ({ form }) => {
    const { register, control, formState } = form;
    const { fields, append, remove } = useFieldArray({ control, name: "attacks" });

    return (
        <div className={styles.step}>
            <h2 className={styles.stepTitle}>Attacks</h2>

            {fields.length === 0 && <p className={styles.hint}>No attacks yet.</p>}

            {fields.map((field, index) => {
                const errors = formState.errors.attacks?.[index];

                return (
                    <fieldset key={field.id} className={styles.attack}>
                        <legend>Attack {index + 1}</legend>

                        <Field id={`attack-name-${index}`} label="Name" error={errors?.name?.message}>
                            <input
                                id={`attack-name-${index}`}
                                type="text"
                                {...register(`attacks.${index}.name`, {
                                    required: "Attack name is required.",
                                    // The API keys attacks by (sheet, name), so duplicates collide.
                                    validate: (value, values) =>
                                        values.attacks.filter((attack) => attack.name === value).length === 1 ||
                                        "Attack names must be unique.",
                                })}
                            />
                        </Field>

                        <Field
                            id={`attack-amount-${index}`}
                            label="Dice amount"
                            error={errors?.attack_dice_amount?.message}
                        >
                            <input
                                id={`attack-amount-${index}`}
                                type="number"
                                min={1}
                                {...register(`attacks.${index}.attack_dice_amount`, {
                                    valueAsNumber: true,
                                    min: { value: 1, message: "Must roll at least one die." },
                                })}
                            />
                        </Field>

                        <Field id={`attack-dice-${index}`} label="Dice">
                            <select id={`attack-dice-${index}`} {...register(`attacks.${index}.attack_dice`)}>
                                {DICE.map((dice) => (
                                    <option key={dice} value={dice}>
                                        {dice}
                                    </option>
                                ))}
                            </select>
                        </Field>

                        <Field id={`attack-damage-${index}`} label="Damage type">
                            <input
                                id={`attack-damage-${index}`}
                                type="text"
                                {...register(`attacks.${index}.damage_type`)}
                            />
                        </Field>

                        <button type="button" className={styles.remove} onClick={() => remove(index)}>
                            Remove attack
                        </button>
                    </fieldset>
                );
            })}

            <button
                type="button"
                className={styles.secondary}
                onClick={() => append({ name: "", attack_dice: "d6", attack_dice_amount: 1, damage_type: "" })}
            >
                Add attack
            </button>
        </div>
    );
};

export default StepAttacks;
