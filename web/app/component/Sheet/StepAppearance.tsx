import type { TSheetForm } from "@/domain/d7d/sheet";
import type { UseFormReturn } from "react-hook-form";

import Field from "./Field";

import styles from "./sheet.module.scss";

const TRAITS = [
    { key: "height", label: "Height" },
    { key: "weight", label: "Weight" },
    { key: "eyes", label: "Eyes" },
    { key: "skin", label: "Skin" },
    { key: "hair", label: "Hair" },
] as const;

const StepAppearance: React.FC<{ form: UseFormReturn<TSheetForm> }> = ({ form }) => {
    const { register, formState } = form;

    return (
        <div className={styles.step}>
            <h2 className={styles.stepTitle}>Appearance & backstory</h2>
            <p className={styles.hint}>Everything here is optional.</p>

            <div className={styles.traits}>
                <Field id="age" label="Age" error={formState.errors.age?.message}>
                    <input
                        id="age"
                        type="number"
                        min={0}
                        max={255}
                        {...register("age", {
                            valueAsNumber: true,
                            min: { value: 0, message: "Age must be between 0 and 255." },
                            max: { value: 255, message: "Age must be between 0 and 255." },
                        })}
                    />
                </Field>

                {TRAITS.map((trait) => (
                    <Field key={trait.key} id={trait.key} label={trait.label}>
                        <input id={trait.key} type="text" {...register(trait.key)} />
                    </Field>
                ))}
            </div>

            <Field id="appearance" label="Appearance">
                <textarea id="appearance" rows={3} {...register("appearance")} />
            </Field>

            <Field id="backstory" label="Backstory">
                <textarea id="backstory" rows={6} {...register("backstory")} />
            </Field>
        </div>
    );
};

export default StepAppearance;
