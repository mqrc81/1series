import { QueryClient as ReactQueryClient } from 'react-query';
import { ApiResponse, create } from 'apisauce';
import { UseQueryOptions } from 'react-query/types/react/types';
import { QueryKey } from 'react-query/types/core/types';

const THIRTY_SECONDS = 30 * 1000;

export const QueryClient = new ReactQueryClient({
    defaultOptions: {
        queries: {
            staleTime: THIRTY_SECONDS,
        },
    },
});

export const Apisauce = create({
    baseURL: import.meta.env.VITE_BACKEND_URL + '/api',
    xsrfCookieName: '_csrf',
    xsrfHeaderName: 'X-CSRF-Token',
});

Apisauce.addResponseTransform((response: ApiResponse<any>) => {
    if (!response.ok) throw response;
});

export type QueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseQueryOptions<TQueryFnData, TError, TData, TQueryKey>, 'queryKey' | 'queryFn'>

export const GetQuery = <TData>(url: string, params = {}): () => Promise<TData> => {
    return async () => {
        const {data} = await Apisauce.get<TData>(url, params);

        return data;
    };
};

export * from './shows/queries';
export * from './shows/dtos';
