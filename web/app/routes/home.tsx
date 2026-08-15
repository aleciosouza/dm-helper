import { requireAuth, requireAuthClient } from "../domain/auth/authMiddleware";
import type { Route } from "./+types/home";
import HomeComponent from "@/component/Home/Home";

export function meta(_: Route.MetaArgs) {
    return [{ title: "DM Helper" }, { name: "description", content: "Welcome to DM Helper!" }];
}

export const middleware: Route.MiddlewareFunction[] = [requireAuth];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireAuthClient];

export default function Home() {
    return <HomeComponent />;
}
