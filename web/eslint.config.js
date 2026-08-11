import eslintReact from "@eslint-react/eslint-plugin";
import js from "@eslint/js";
import prettierCompat from "eslint-config-prettier/flat";
import jsxA11y from "eslint-plugin-jsx-a11y";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import { defineConfig, globalIgnores } from "eslint/config";
import globals from "globals";
import tseslint from "typescript-eslint";

/**
 * Exceptions to the "react-refresh/only-export-components" rule,
 * which is used in React Router route modules.
 */
const ROUTE_MODULE_EXPORTS = [
    "meta",
    "links",
    "headers",
    "handle",
    "shouldRevalidate",
    "loader",
    "clientLoader",
    "clientLoaderHydrate",
    "action",
    "clientAction",
    "middleware",
    "clientMiddleware",
    "HydrateFallback",
    "ErrorBoundary",
    "Layout",
];

/**
 * Disables eslint-react rules already covered by eslint-plugin-react-hooks.
 */
const ESLINT_REACT_HOOK_DUPLICATES = {
    "@eslint-react/error-boundaries": "off",
    "@eslint-react/exhaustive-deps": "off",
    "@eslint-react/purity": "off",
    "@eslint-react/rules-of-hooks": "off",
    "@eslint-react/set-state-in-effect": "off",
    "@eslint-react/set-state-in-render": "off",
    "@eslint-react/static-components": "off",
    "@eslint-react/unsupported-syntax": "off",
    "@eslint-react/use-memo": "off",
    "@eslint-react/use-state": "off",
};

export default defineConfig(
    globalIgnores(["build/", ".react-router/", "public/"]),
    {
        name: "dm-helper/linter-options",
        linterOptions: { reportUnusedDisableDirectives: "error" },
    },

    {
        name: "dm-helper/typescript",
        files: ["**/*.{ts,tsx,mts}"],
        extends: [js.configs.recommended, tseslint.configs.recommendedTypeChecked],
        languageOptions: {
            parserOptions: {
                projectService: true,
                tsconfigRootDir: import.meta.dirname,
            },
        },
        rules: {
            "@typescript-eslint/consistent-type-imports": "error",
            "@typescript-eslint/no-import-type-side-effects": "error",

            // Permite `_arg` e `catch (_err)`
            "@typescript-eslint/no-unused-vars": [
                "error",
                {
                    args: "all",
                    argsIgnorePattern: "^_",
                    varsIgnorePattern: "^_",
                    caughtErrorsIgnorePattern: "^_",
                    destructuredArrayIgnorePattern: "^_",
                    ignoreRestSiblings: true,
                },
            ],

            // Retricts deprecated imports from React Router 7+
            "@typescript-eslint/no-restricted-imports": [
                "error",
                {
                    paths: [
                        {
                            name: "react-router-dom",
                            message: "Deprecated on React Router 7+, use 'react-router' instead.",
                        },
                        {
                            name: "@remix-run/react",
                            message: "Deprecated on React Router 7+, use 'react-router' instead.",
                        },
                    ],
                },
            ],
        },
    },

    // React
    {
        name: "dm-helper/react",
        files: ["**/*.{ts,tsx}"],
        extends: [
            eslintReact.configs["recommended-type-checked"],
            reactHooks.configs.flat["recommended-latest"],
            jsxA11y.flatConfigs.recommended,
        ],
        rules: ESLINT_REACT_HOOK_DUPLICATES,
    },

    // Route modules
    {
        name: "dm-helper/react-router-routes",
        files: ["app/**/*.tsx"],
        extends: [reactRefresh.configs.vite],
        rules: {
            "react-refresh/only-export-components": [
                "warn",
                {
                    allowConstantExport: true,
                    allowExportNames: ROUTE_MODULE_EXPORTS,
                },
            ],
        },
    },

    {
        name: "dm-helper/browser-env",
        files: ["app/**/*.{ts,tsx}"],
        languageOptions: { globals: globals.browser },
    },
    {
        name: "dm-helper/server-env",
        files: [
            "app/routes.ts",
            "app/**/entry.server.{ts,tsx}",
            "app/**/*.server.{ts,tsx}",
            "app/**/.server/**/*.{ts,tsx}",
            "*.config.{ts,js}",
        ],
        languageOptions: { globals: globals.node },
    },

    // Exclude JS files from TypeScript rules
    {
        name: "dm-helper/js-config-files",
        files: ["**/*.{js,mjs,cjs}"],
        extends: [js.configs.recommended, tseslint.configs.disableTypeChecked],
        languageOptions: { globals: globals.node },
    },

    prettierCompat,
);
