interface ImportMetaEnv {
    readonly VITE_API_URL: string;
    /** `true` in server-side, `false` in client-side */
    readonly SSR: boolean;
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}
