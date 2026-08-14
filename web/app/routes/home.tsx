import { useNavigate } from "react-router";
import { useQueryClient } from "@tanstack/react-query";

import { requireAuth, requireAuthClient } from "../domain/auth/authMiddleware";
import { useUser } from "../domain/auth/authQuery";
import { clearToken } from "../domain/auth/token";
import type { Route } from "./+types/home";

export function meta(_: Route.MetaArgs) {
    return [{ title: "DM Helper" }, { name: "description", content: "Welcome to DM Helper!" }];
}

export const middleware: Route.MiddlewareFunction[] = [requireAuth];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireAuthClient];

export default function Home() {
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
