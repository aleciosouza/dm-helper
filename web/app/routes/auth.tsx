import AuthComponent from "@/component/Auth";

import { requireGuest, requireGuestClient } from "../domain/auth/authMiddleware";
import type { Route } from "./+types/auth";

export function meta(_: Route.MetaArgs) {
    return [{ title: "Login · DM Helper" }];
}

export const middleware: Route.MiddlewareFunction[] = [requireGuest];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireGuestClient];

export default function Auth() {
    return <AuthComponent />;
}
