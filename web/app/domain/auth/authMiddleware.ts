import { redirect, type DataStrategyResult, type MiddlewareFunction } from "react-router";

import apiFetch, { isApiError } from "../api/api";
import { clearToken, expiredTokenCookie, getToken } from "./token";

const LOGIN_ROUTE = "/auth";
const HOME_ROUTE = "/home";

type ServerMiddleware = MiddlewareFunction<Response>;
type ClientMiddleware = MiddlewareFunction<Record<string, DataStrategyResult>>;

async function isAuthenticated(request?: Request): Promise<boolean> {
    if (!getToken(request)) return false;

    try {
        await apiFetch<void>("auth", { method: "HEAD", request });
        return true;
    } catch (error) {
        if (isApiError(error) && error.status === 401) return false;
        throw error;
    }
}

export const requireAuth: ServerMiddleware = async ({ request }) => {
    if (await isAuthenticated(request)) return;

    throw redirect(LOGIN_ROUTE, { headers: { "Set-Cookie": expiredTokenCookie() } });
};

export const requireAuthClient: ClientMiddleware = async () => {
    if (await isAuthenticated()) return;

    clearToken();
    throw redirect(LOGIN_ROUTE);
};

export const requireGuest: ServerMiddleware = async ({ request }) => {
    if (await isAuthenticated(request).catch(() => false)) {
        throw redirect(HOME_ROUTE);
    }
};

export const requireGuestClient: ClientMiddleware = async () => {
    if (await isAuthenticated().catch(() => false)) throw redirect(HOME_ROUTE);
};
