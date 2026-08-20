import { useState } from "react";
import { useNavigate } from "react-router";
import { useForm, type DefaultValues, type FieldPath } from "react-hook-form";

import { isApiError } from "@/domain/api/api";
import type { TSheetForm } from "@/domain/d7d/sheet";
import { useCreateSheet, useUpdateSheet, toSheetPayload } from "@/domain/d7d/sheet";

import StepAppearance from "./StepAppearance";
import StepAttacks from "./StepAttacks";
import StepIdentity from "./StepIdentity";
import Stepper from "./Stepper";
import StepPersonality from "./StepPersonality";
import StepSkills from "./StepSkills";

import styles from "./sheet.module.scss";

interface Step {
    title: string;
    fields: FieldPath<TSheetForm>[];
    render: (form: ReturnType<typeof useForm<TSheetForm>>) => React.ReactNode;
}

const STEPS: readonly Step[] = [
    {
        title: "Identity",
        fields: [
            "name",
            "class",
            "race",
            "alignment",
            "level",
            "str",
            "dex",
            "con",
            "int",
            "wis",
            "cha",
            "save_str",
            "save_dex",
            "save_con",
            "save_int",
            "save_wis",
            "save_cha",
        ],
        render: (form) => <StepIdentity form={form} />,
    },
    {
        title: "Skills",
        fields: ["skills"],
        render: (form) => <StepSkills form={form} />,
    },
    {
        title: "Attacks",
        fields: ["attacks"],
        render: (form) => <StepAttacks form={form} />,
    },
    {
        title: "Personality",
        fields: ["background", "personality_traits", "ideals", "bonds", "flaws"],
        render: (form) => <StepPersonality form={form} />,
    },
    {
        title: "Appearance",
        fields: ["age", "height", "weight", "eyes", "skin", "hair", "appearance", "backstory"],
        render: (form) => <StepAppearance form={form} />,
    },
];

interface ISheetFormProps {
    sheetId?: number;
    defaultValues: DefaultValues<TSheetForm>;
}

const SheetForm: React.FC<ISheetFormProps> = ({ sheetId, defaultValues }) => {
    const [step, setStep] = useState(0);
    const [error, setError] = useState<string | null>(null);
    const form = useForm<TSheetForm>({ defaultValues, mode: "onChange" });

    const createSheet = useCreateSheet();
    const updateSheet = useUpdateSheet();

    const navigate = useNavigate();

    const isLastStep = step === STEPS.length - 1;
    const currentStep = STEPS[step];

    async function handleNext() {
        if (await form.trigger(currentStep.fields, { shouldFocus: true })) {
            setStep(step + 1);
        }
    }

    async function handleSubmit(values: TSheetForm) {
        setError(null);

        const payload = toSheetPayload(values);

        try {
            if (sheetId === undefined) {
                await createSheet.mutateAsync(payload);
            } else {
                await updateSheet.mutateAsync({ ...payload, id: sheetId });
            }

            await navigate("/home");
        } catch (submitError) {
            const message = isApiError(submitError) && typeof submitError.data === "string" ? submitError.data : null;
            setError(message ?? "Could not save the sheet. Please try again.");
        }
    }

    return (
        <main className={styles.page}>
            <h1 className={styles.title}>{sheetId === undefined ? "New character sheet" : "Edit character sheet"}</h1>

            <Stepper step={step} total={STEPS.length} title={currentStep.title} />

            <form className={styles.form} onSubmit={(event) => void form.handleSubmit(handleSubmit)(event)}>
                {error && (
                    <p className={styles.error} role="alert">{error}</p>
                )}

                {currentStep.render(form)}

                <div className={styles.actions}>
                    <button
                        type="button"
                        className={styles.secondary}
                        disabled={step === 0 || form.formState.isSubmitting}
                        onClick={() => setStep(step - 1)}
                    >
                        Back
                    </button>

                    {isLastStep ? (
                        <button
                            type="submit"
                            className={styles.submit}
                            disabled={!form.formState.isValid || form.formState.isSubmitting}
                        >
                            {form.formState.isSubmitting ? "Saving..." : "Save sheet"}
                        </button>
                    ) : (
                        <button
                            type="button"
                            className={styles.submit}
                            disabled={!form.formState.isValid}
                            onClick={() => void handleNext()}
                        >
                            Next
                        </button>
                    )}
                </div>
            </form>
        </main>
    );
};

export default SheetForm;
