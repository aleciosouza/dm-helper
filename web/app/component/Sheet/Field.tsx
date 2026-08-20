import styles from "./sheet.module.scss";

interface FieldProps {
    id: string;
    label: string;
    error?: string;
    children: React.ReactNode;
}

const Field: React.FC<FieldProps> = ({ id, label, error, children }) => {
    return (
        <div className={styles.field}>
            <label htmlFor={id}>{label}</label>
            {children}
            {error && (
                <p className={styles.error} role="alert">
                    {error}
                </p>
            )}
        </div>
    );
};

export default Field;
