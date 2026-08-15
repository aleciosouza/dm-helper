import { useState } from "react";

import styles from "./auth.module.scss";
import { useLogin } from "@/domain/auth/authQuery";
import { useNavigate } from "react-router";

const Auth: React.FC = () => {
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
        <main className={styles.page}>
            <div className={styles.card}>
                <h1 className={styles.title}>Login</h1>

                {login.isError && (
                    <p className={styles.error} role="alert">
                        Invalid email or password.
                    </p>
                )}

                <form className={styles.form} onSubmit={(event) => void handleSubmit(event)}>
                    <div className={styles.field}>
                        <label htmlFor="email">E-mail</label>
                        <input
                            id="email"
                            type="email"
                            autoComplete="email"
                            required
                            value={email}
                            onChange={(event) => setEmail(event.target.value)}
                        />
                    </div>

                    <div className={styles.field}>
                        <label htmlFor="password">Password</label>
                        <input
                            id="password"
                            type="password"
                            autoComplete="current-password"
                            required
                            value={password}
                            onChange={(event) => setPassword(event.target.value)}
                        />
                    </div>

                    <button className={styles.submit} type="submit" disabled={login.isPending}>
                        {login.isPending ? "Logging in..." : "Login"}
                    </button>
                </form>
            </div>
        </main>
    );
};

export default Auth;
