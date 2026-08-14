import { useState } from "react";
import { useNavigate } from "react-router";

import { requireGuest, requireGuestClient } from "../domain/auth/authMiddleware";
import { useLogin } from "../domain/auth/authQuery";
import type { Route } from "./+types/auth";

export function meta(_: Route.MetaArgs) {
    return [{ title: "Entrar · DM Helper" }];
}

export const middleware: Route.MiddlewareFunction[] = [requireGuest];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireGuestClient];

export default function Auth() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const login = useLogin();
    const navigate = useNavigate();

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        await login.mutateAsync({ email, password });
        await navigate("/home");
    }

    return (
        <main>
            <h1>Entrar</h1>
            <form onSubmit={(event) => void handleSubmit(event)}>
                <label htmlFor="email">E-mail</label>
                <input
                    id="email"
                    type="email"
                    autoComplete="email"
                    required
                    value={email}
                    onChange={(event) => setEmail(event.target.value)}
                />

                <label htmlFor="password">Senha</label>
                <input
                    id="password"
                    type="password"
                    autoComplete="current-password"
                    required
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                />

                <button type="submit" disabled={login.isPending}>
                    {login.isPending ? "Entrando..." : "Entrar"}
                </button>

                {login.isError && <p role="alert">E-mail ou senha inválidos.</p>}
            </form>
        </main>
    );
}
