
import type { QueryClient } from "@tanstack/react-query";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import apiFetch from "@/domain/api";

import type { ISheet } from "./sheet";
import type { TSheetForm } from "./sheetForm";

const SHEET_QUERY_KEY = ["sheet", "get"];

export function useSheets() {
    return useQuery({
        queryKey: SHEET_QUERY_KEY,
        queryFn: () => apiFetch<ISheet[]>("sheet")
    })
}

export function useSheet(id?: number) {
    return useQuery({
        queryKey: [...SHEET_QUERY_KEY, id],
        queryFn: () => apiFetch<ISheet>(`sheet/${id}`),
        enabled: id !== undefined,
    })
}

/**
 * Ensures that the "getAll" query never returns a stale sheet
 */
function onSheetSaved(queryClient: QueryClient, sheet: ISheet) {
    queryClient.setQueryData([...SHEET_QUERY_KEY, sheet.id], sheet);
    void queryClient.invalidateQueries({ queryKey: SHEET_QUERY_KEY, exact: true });
}

export function useCreateSheet() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (sheet: TSheetForm) =>
            apiFetch<ISheet>("sheet", { method: "POST", body: JSON.stringify(sheet) }),
        onSuccess: (sheet) => onSheetSaved(queryClient, sheet),
    });
}

export function useUpdateSheet() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (sheet: ISheet) =>
            apiFetch<ISheet>(`sheet/${sheet.id}`, { method: "PATCH", body: JSON.stringify(sheet) }),
        onSuccess: (sheet) => onSheetSaved(queryClient, sheet),
    });
}

export function useDeleteSheet() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (id: number) => apiFetch<ISheet>(`sheet/${id}`, { method: "DELETE" }),
        onSuccess: (sheet) => queryClient.removeQueries({
            queryKey: [...SHEET_QUERY_KEY, sheet.id],
            exact: true,
        })
    });
}
