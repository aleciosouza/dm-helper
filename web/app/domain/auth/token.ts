const TOKEN_KEY = "dmh_t";
const TOKEN_MAX_AGE = 60 * 60 * 24 * 7;

function serializeCookie(value: string, maxAge: number): string {
    const secure = typeof location !== "undefined" && location.protocol === "https:" ? "; Secure" : "";
    return `${TOKEN_KEY}=${value}; Path=/; Max-Age=${maxAge}; SameSite=Lax${secure}`;
}

function readCookie(cookieHeader: string | null | undefined): string | null {
    if (!cookieHeader) return null;

    for (const pair of cookieHeader.split(";")) {
        const separator = pair.indexOf("=");
        if (separator === -1) continue;

        if (pair.slice(0, separator).trim() === TOKEN_KEY) {
            return decodeURIComponent(pair.slice(separator + 1).trim()) || null;
        }
    }

    return null;
}

export function getToken(request?: Request): string | null {
    if (request) {
        return readCookie(request.headers.get("Cookie"));
    }


    if (typeof document === "undefined") return null;

    return readCookie(document.cookie);
}

export function setToken(token: string): void {
    document.cookie = serializeCookie(encodeURIComponent(token), TOKEN_MAX_AGE);
}

export function clearToken(): void {
    document.cookie = expiredTokenCookie();
}

export function expiredTokenCookie(): string {
    return serializeCookie("", 0);
}
