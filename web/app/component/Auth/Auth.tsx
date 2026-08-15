import { useState } from "react";

import styles from "./auth.module.scss";

import Login from "./Login";
import Register from "./Register";

const Auth: React.FC = () => {
    const [isLogin, setIsLogin] = useState(true);

    return (
        <div className={styles.container}>
            <main className={styles.page}>
                <div className={styles.card}>
                    <h1 className={styles.title}>{isLogin ? "Login" : "Register"}</h1>
                    {isLogin ? <Login /> : <Register />}

                    <button className={styles.anchor} onClick={() => setIsLogin(!isLogin)}>
                        {isLogin ? (
                            <span>Don't have an account? <b>Register</b></span>
                        ) : (
                            <span>Already have an account? <b>Login</b></span>
                        )}
                    </button>
                </div>
            </main>
        </div>
    );


};

export default Auth;
