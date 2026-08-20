import type { TSheetForm } from "@/domain/d7d/sheet";
import type { UseFormReturn } from "react-hook-form";

import Field from "./Field";

import styles from "./sheet.module.scss";

const TRAITS = [
    { key: "personality_traits", label: "Personality traits" },
    { key: "ideals", label: "Ideals" },
    { key: "bonds", label: "Bonds" },
    { key: "flaws", label: "Flaws" },
] as const;

const StepPersonality: React.FC<{ form: UseFormReturn<TSheetForm> }> = ({ form }) => {
    return (
        <div className={styles.step}>
            <h2 className={styles.stepTitle}>Personality</h2>
            <p className={styles.hint}>Everything here is optional.</p>

            <Field id="background" label="Background">
                <input id="background" type="text" {...form.register("background")} />
            </Field>

            {TRAITS.map((trait) => (
                <Field key={trait.key} id={trait.key} label={trait.label}>
                    <textarea id={trait.key} rows={3} {...form.register(trait.key)} />
                </Field>
            ))}
        </div>
    );
};

export default StepPersonality;
