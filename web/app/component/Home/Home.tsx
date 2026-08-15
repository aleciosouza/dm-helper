import { useUser } from '@/domain/auth/authQuery';
import { clearToken } from '@/domain/auth/token';
import { useQueryClient } from '@tanstack/react-query';
import { useNavigate } from 'react-router';

const Home = () => {
    const { data: user, isPending } = useUser();
    const queryClient = useQueryClient();
    const navigate = useNavigate();

    async function handleLogout() {
        clearToken();
        queryClient.clear();
        await navigate("/auth");
    }

    return (
        <main>
            <h1>DM Helper</h1>

            <p>{isPending ? "Carregando..." : `Olá, ${user?.name || user?.email}`}</p>

            <button type="button" onClick={() => void handleLogout()}>
                Sair
            </button>
        </main>
    );

}

export default Home
