import styles from "./sheet.module.scss";

interface StepperProps {
    step: number;
    total: number;
    title: string;
}

const Stepper: React.FC<StepperProps> = ({ step, total, title }) => {
    return (
        <div className={styles.stepper}>
            <p className={styles.stepperCount}>
                Step {step + 1} of {total} — {title}
            </p>

            <div
                className={styles.progress}
                role="progressbar"
                aria-valuenow={step + 1}
                aria-valuemin={1}
                aria-valuemax={total}
                aria-label="Sheet progress"
            >
                <span className={styles.progressBar} style={{ inlineSize: `${((step + 1) / total) * 100}%` }} />
            </div>
        </div>
    );
};

export default Stepper;
