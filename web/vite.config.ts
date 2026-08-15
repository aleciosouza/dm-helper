import { fileURLToPath } from "node:url";

import { reactRouter } from "@react-router/dev/vite";
import { defineConfig } from "vite";

const stylesDir = fileURLToPath(new URL("./app/styles", import.meta.url));

export default defineConfig({
    plugins: [reactRouter()],
    css: {
        preprocessorOptions: {
            scss: {
                // Allow `@use "abstracts" as *;`
                loadPaths: [stylesDir]
            }
        }
    },
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
