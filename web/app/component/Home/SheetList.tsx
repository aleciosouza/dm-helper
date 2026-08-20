import type { ISheet } from "@/domain/d7d/sheet";
import { useSheets } from "@/domain/d7d/sheet";
import { useNavigate } from "react-router";

import styles from "./home.module.scss";


const Sheet: React.FC<{ sheet: ISheet }> = ({ sheet }) => {
    const navigate = useNavigate();

    async function onClickEdit() {
        await navigate(`/sheet/${sheet.id}`);
        return
    }

    return (
        <div className={styles.sheet}>
            <p>{sheet.name}</p>
            <p>{sheet.race}, {sheet.class}, lvl {sheet.level} | <b>{sheet.alignment}</b></p>
            <button type="button" name='edit' onClick={() => void onClickEdit()}>Edit</button>
        </div>
    )
}

const SheetsList: React.FC = () => {
    const { data: sheets, isPending } = useSheets();
    const navigate = useNavigate();

    async function onClickCreate() {
        await navigate("/sheet");
    }

    if (isPending || !sheets) {
        return (<div><p>Loading sheets...</p></div>);
    }

    if (sheets.length === 0) {
        return (
            <div className={styles.sheets}>
                <div className={styles.sheet}>
                    <h1>Create a character sheet and get your adventure started!</h1>
                    <button type="button" onClick={() => void onClickCreate()}>+</button>
                </div>
            </div>
        );
    }

    return (
        <div className={styles.sheets}>
            {sheets.map(sheet => <Sheet key={sheet.id} sheet={sheet} />)}
            <button type="button" onClick={() => void onClickCreate()}>New sheet</button>
        </div>
    );
}

export default SheetsList;
