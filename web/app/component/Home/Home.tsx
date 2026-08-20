import { useUser } from '@/domain/auth/authQuery';
import { useNavigate } from 'react-router';

import styles from './home.module.scss';
import SheetsList from './SheetList';

const ProfileBar: React.FC = () => {
    const { data: user, isPending } = useUser();
    const navigate = useNavigate();

    const handleLogout = async () => {
        await navigate("/auth");
    }

    return (
        <div className={styles.header}>
            <span>{isPending ? "Carregando..." : `Olá, ${user?.name || user?.email}`}</span>

            <button type="button" className={styles.header__logout} onClick={() => void handleLogout()}>
                Logout
            </button>
        </div>
    );
}

const Home: React.FC = () => {
    return (
        <>
            <ProfileBar />
            <SheetsList />
        </>
    );

}

export default Home
