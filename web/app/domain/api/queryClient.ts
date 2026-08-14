import { QueryClient } from "@tanstack/react-query";

function makeQueryClient(): QueryClient {
    return new QueryClient({
        defaultOptions: {
            queries: {
                staleTime: 30_000,
                retry: (failureCount, error) => {
                    if ("status" in error && (error as { status: number }).status < 500) {
                        return false;
                    }

                    return failureCount < 2;
                },
            },
        },
    });
}

let browserQueryClient: QueryClient | undefined;

export function getQueryClient(): QueryClient {
    if (typeof window === "undefined") return makeQueryClient();
    browserQueryClient ??= makeQueryClient();
    return browserQueryClient;
}
