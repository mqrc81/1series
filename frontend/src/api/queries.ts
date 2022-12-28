import { QueryFunctionContext, QueryKey } from 'react-query/types/core/types';
import { UseQueryOptions } from 'react-query/types/react/types';
import { GetNextPageParamFunction, GetPreviousPageParamFunction, UseInfiniteQueryOptions } from 'react-query';
import { Paginated } from './dtos';
import { ApisauceClient } from '../providers';

export type QueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseQueryOptions<TQueryFnData, TError, TData, TQueryKey>, 'queryKey' | 'queryFn'>
export type InfiniteQueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseInfiniteQueryOptions<TQueryFnData, TError, TData, TQueryFnData, TQueryKey>, 'queryKey'>

export const getPreviousPageParam: GetPreviousPageParamFunction<Paginated<unknown>> = ({previousPage = undefined}) => previousPage;
export const getNextPageParam: GetNextPageParamFunction<Paginated<unknown>> = ({nextPage = undefined}) => nextPage;

export const GetQuery = <TData>(url: string, params = {}): () => Promise<TData> => {
    return async () => {
        const {data} = await ApisauceClient.get<TData>(url, params);

        return data;
    };
};

export const GetInfiniteQuery = <TData>(url: string): (context: QueryFunctionContext) => Promise<TData> => {
    return async ({pageParam = 1}) => {
        const {data} = await ApisauceClient.get<TData>(url, {page: pageParam});

        return data;
    };
};
