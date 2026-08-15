import { useState } from "react";
import { useNavigate } from "react-router";

import styles from "./auth.module.scss";
import { useRegister } from "@/domain/auth/authQuery";

const Register: React.FC = () => {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState<string | null>(null);

    const register = useRegister();
    const navigate = useNavigate();

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        setError(null);

        if (password !== confirmPassword) {
            setError("Passwords do not match.");
            return;
        }

        await register.mutateAsync({ name, email, password });
        await navigate("/home");
    }

    return (
        <form className={styles.form} onSubmit={(event) => void handleSubmit(event)}>
            {error && (
                <p className={styles.error} role="alert">
                    {error}
                </p>
            )}

            <div className={styles.field}>
                <label htmlFor="name">Name</label>
                <input
                    id="name"
                    type="text"
                    autoComplete="name"
                    required
                    value={name}
                    onChange={(event) => setName(event.target.value)}
                />
            </div>

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
                    autoComplete="new-password"
                    required
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                />
            </div>

            <div className={styles.field}>
                <label htmlFor="confirm-password">Confirm Password</label>
                <input
                    id="confirm-password"
                    type="password"
                    autoComplete="new-password"
                    required
                    value={confirmPassword}
                    onChange={(event) => setConfirmPassword(event.target.value)}
                />
            </div>

            <button className={styles.submit} type="submit">
                Register
            </button>
        </form>
    )
}

export default Register;
