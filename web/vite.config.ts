import { reactRouter } from "@react-router/dev/vite";
import { defineConfig } from "vite";

export default defineConfig({
    plugins: [reactRouter()],
    resolve: {
        tsconfigPaths: true,
    },
    server: {
        proxy: {
            "/api": {
                // @TODO: Replace before deploying to production
                target: "http://localhost:8080",
                changeOrigin: true
            }
        }
    }
});
