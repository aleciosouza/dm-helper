import { useNavigate } from "react-router";
import Auth from "@/component/Auth";

import { requireGuest, requireGuestClient } from "../domain/auth/authMiddleware";
import type { Route } from "./+types/auth";

export function meta(_: Route.MetaArgs) {
    return [{ title: "Login · DM Helper" }];
}

export const middleware: Route.MiddlewareFunction[] = [requireGuest];
export const clientMiddleware: Route.ClientMiddlewareFunction[] = [requireGuestClient];

export default function AuthRoute() {
    const navigate = useNavigate();

    const handleOnSubmit = async () => await navigate("/home");

    return <Auth onSubmit={() => void handleOnSubmit()} />;
}
