import { useSheet, defaultSheetForm, toSheetForm } from "@/domain/d7d/sheet";

import SheetForm from "./SheetForm";

import styles from "./sheet.module.scss";

const Sheet: React.FC<{ id?: string }> = ({ id }) => {
    const sheetId = id === undefined ? undefined : Number(id);

    const { data: sheet, isPending, isError } = useSheet(sheetId);

    if (sheetId === undefined) {
        return <SheetForm defaultValues={defaultSheetForm()} />;
    }

    if (Number.isNaN(sheetId)) {
        return (
            <p className={styles.notice} role="alert">
                Sheet not found.
            </p>
        );
    }

    if (isPending) {
        return <p className={styles.notice}>Loading sheet...</p>;
    }

    if (isError || !sheet) {
        return (
            <p className={styles.notice} role="alert">
                Sheet not found.
            </p>
        );
    }

    return <SheetForm sheetId={sheet.id} defaultValues={toSheetForm(sheet)} />;
};

export default Sheet;
