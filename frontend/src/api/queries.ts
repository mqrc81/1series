import { QueryFunctionContext, QueryKey } from 'react-query/types/core/types';
import { UseQueryOptions } from 'react-query/types/react/types';
import { UseInfiniteQueryOptions } from 'react-query';
import { ApiClient } from './index';

export type QueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseQueryOptions<TQueryFnData, TError, TData, TQueryKey>, 'queryKey' | 'queryFn'>
export type InfiniteQueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseInfiniteQueryOptions<TQueryFnData, TError, TData, TQueryFnData, TQueryKey>, 'queryKey'>

export const GetQuery = <TData>(url: string, params = {}): (params: {}) => Promise<TData> => {
    return async () => {
        const {data} = await ApiClient.get<TData>(url, params);

        return data;
    };
};

export const GetInfiniteQuery = <TData>(url: string): (context: QueryFunctionContext) => Promise<TData> => {
    return async ({pageParam}) => {
        const {data} = await ApiClient.get<TData>(url, {page: pageParam});

        return data;
    };
};
