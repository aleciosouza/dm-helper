import { ABILITIES, ALIGNMENTS, CLASSES, RACES } from "@/domain/d7d/options";
import type { TSheetForm } from "@/domain/d7d/sheet";
import type { UseFormReturn } from "react-hook-form";

import Field from "./Field";

import styles from "./sheet.module.scss";

const StepIdentity: React.FC<{ form: UseFormReturn<TSheetForm> }> = ({ form }) => {
    const { register, formState } = form;
    const { errors } = formState;

    return (
        <div className={styles.step}>
            <Field id="name" label="Name" error={errors?.name?.message}>
                <input id="name" type="text" {...register("name", { required: "Name is required." })} />
            </Field>

            <Field id="class" label="Class" error={errors?.class?.message}>
                <select id="class" {...register("class", { required: "Class is required." })}>
                    <option value="">Select a class</option>
                    {CLASSES.map((option) => (
                        <option key={option.value} value={option.value}>
                            {option.value}
                        </option>
                    ))}
                </select>
            </Field>

            <Field id="race" label="Race" error={errors?.race?.message}>
                <select id="race" {...register("race", { required: "Race is required." })}>
                    <option value="">Select a race</option>
                    {RACES.map((option) => (
                        <option key={option.value} value={option.value}>
                            {option.value}
                        </option>
                    ))}
                </select>
            </Field>

            <Field id="alignment" label="Alignment (optional)">
                <select id="alignment" {...register("alignment")}>
                    <option value="">Select an alignment</option>
                    {ALIGNMENTS.map((alignment) => (
                        <option key={alignment} value={alignment}>
                            {alignment}
                        </option>
                    ))}
                </select>
            </Field>

            <Field id="level" label="Level" error={errors?.level?.message}>
                <input
                    id="level"
                    type="number"
                    min={1}
                    max={20}
                    {...register("level", {
                        required: "Level is required.",
                        valueAsNumber: true,
                        min: { value: 1, message: "Level must be between 1 and 20." },
                        max: { value: 20, message: "Level must be between 1 and 20." },
                    })}
                />
            </Field>

            <fieldset className={styles.abilities}>
                <legend>Ability scores</legend>

                {ABILITIES.map((ability) => (
                    <div key={ability.key} className={styles.ability}>
                        <Field id={ability.key} label={ability.label} error={errors[ability.key]?.message}>
                            <input
                                id={ability.key}
                                type="number"
                                min={1}
                                max={30}
                                {...register(ability.key, {
                                    required: `${ability.label} is required.`,
                                    valueAsNumber: true,
                                    min: { value: 1, message: "Must be between 1 and 30." },
                                    max: { value: 30, message: "Must be between 1 and 30." },
                                })}
                            />
                        </Field>

                        <label className={styles.toggle} htmlFor={ability.save}>
                            <input id={ability.save} type="checkbox" {...register(ability.save)} />
                            Saving throw proficiency
                        </label>
                    </div>
                ))}
            </fieldset>
        </div>
    );
};

export default StepIdentity;
