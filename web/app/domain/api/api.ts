import axios, { type AxiosInstance } from "axios";

import { getToken } from "../auth/token";

interface ApiEnvelope<T> {
    message: string;
    data: T;
}

function getBaseUrl(): string {
    if (!import.meta.env.SSR) return import.meta.env.VITE_API_URL;

    return process.env.API_URL ?? "http://localhost:8080/api/v1";
}

function getApi(request?: Request): AxiosInstance {
    const token = getToken(request);

    return axios.create({
        baseURL: getBaseUrl(),
        timeout: 10000,
        headers: {
            "Content-Type": "application/json",
            ...(token ? { Authorization: `Bearer ${token}` } : {}),
        },
    });
}

export interface ApiError extends Error {
    status: number;
    data?: unknown;
}

export function isApiError(error: unknown): error is ApiError {
    return error instanceof Error && "status" in error;
}

export interface ApiFetchOptions extends RequestInit {
    /** Só em código de servidor: repassa o cookie da requisição original. */
    request?: Request;
}

export default async function apiFetch<T>(path: string, options?: ApiFetchOptions): Promise<T> {
    const axiosInstance = getApi(options?.request);

    try {
        const response = await axiosInstance.request<ApiEnvelope<T>>({
            url: path,
            method: options?.method || "GET",
            headers: options?.headers ? Object.fromEntries(new Headers(options.headers).entries()) : undefined,
            data: options?.body,
        });

        return response.data?.data;
    } catch (error) {
        if (axios.isAxiosError(error) && error.response) {
            const apiError = new Error(error.message) as ApiError;
            apiError.status = error.response.status;
            apiError.data = error.response.data;
            throw apiError;
        }

        throw error;
    }
}
