import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import apiFetch from "../api/api";
import { setToken } from "./token";
import type { IUser } from "./user";

export interface ICredentials {
    email: string;
    password: string;
}

export interface IRegisterData extends ICredentials {
    name: string;
}

interface ISession {
    token: string;
    user: IUser;
}

const USER_QUERY_KEY = ["auth", "getMe"];

export function useUser() {
    return useQuery({
        queryKey: USER_QUERY_KEY,
        queryFn: () => apiFetch<IUser>("auth"),
    });
}

function useSessionMutation<T>(path: string) {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (variables: T) =>
            apiFetch<ISession>(path, { method: "POST", body: JSON.stringify(variables) }),
        onSuccess: ({ token, user }) => {
            setToken(token);
            queryClient.setQueryData(USER_QUERY_KEY, user);
        },
    });
}

export function useLogin() {
    return useSessionMutation<ICredentials>("auth/login");
}

export function useRegister() {
    return useSessionMutation<IRegisterData>("auth/register");
}
