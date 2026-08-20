import SheetComponent from "@/component/Sheet";

import { requireAuth, requireAuthClient } from "../domain/auth/authMiddleware";
import type { Route } from "./+types/sheet";

export function meta(_: Route.MetaArgs) {
    return [{ title: "Sheet · DM Helper" }, { name: "description", content: "Create and edit a character sheet." }];
}

export const middleware: Route.MiddlewareFunction[] = [requireAuth];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireAuthClient];

export default function Sheet({ params }: Route.ComponentProps) {
    return <SheetComponent id={params.id} />;
}
