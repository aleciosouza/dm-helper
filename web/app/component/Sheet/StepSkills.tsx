import { SKILLS } from "@/domain/d7d/options";
import type { TSheetForm } from "@/domain/d7d/sheet";
import { Controller, type UseFormReturn } from "react-hook-form";

import styles from "./sheet.module.scss";

/**
 * `skills` is a list of numeric ids, so it goes through a Controller — a plain
 * `register` on grouped checkboxes would hand back strings.
 */
const StepSkills: React.FC<{ form: UseFormReturn<TSheetForm> }> = ({ form }) => {
    return (
        <div className={styles.step}>
            <Controller
                control={form.control}
                name="skills"
                render={({ field }) => (
                    <fieldset className={styles.skills}>
                        <legend>Skill proficiencies</legend>

                        {SKILLS.map((skill) => {
                            const isSelected = field.value.includes(skill.value);

                            return (
                                <label key={skill.value} className={styles.toggle}>
                                    <input
                                        type="checkbox"
                                        checked={isSelected}
                                        onBlur={field.onBlur}
                                        onChange={() =>
                                            field.onChange(
                                                isSelected
                                                    ? field.value.filter((value) => value !== skill.value)
                                                    : [...field.value, skill.value],
                                            )
                                        }
                                    />
                                    {skill.label}
                                </label>
                            );
                        })}
                    </fieldset>
                )}
            />
        </div>
    );
};

export default StepSkills;
