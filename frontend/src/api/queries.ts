import { QueryKey } from 'react-query/types/core/types';
import { UseQueryOptions } from 'react-query/types/react/types';
import { GetNextPageParamFunction, GetPreviousPageParamFunction, UseInfiniteQueryOptions, useQuery } from 'react-query';
import { Paginated } from './dtos';
import { ApisauceClient } from '../providers/apisauce';

export type QueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseQueryOptions<TQueryFnData, TError, TData, TQueryKey>, 'queryKey' | 'queryFn'>
export type InfiniteQueryOptions<TQueryFnData = unknown, TError = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseInfiniteQueryOptions<TQueryFnData, TError, TData, TQueryFnData, TQueryKey>, 'queryKey'>

export const getPreviousPageParam: GetPreviousPageParamFunction<Paginated<unknown>> = ({previousPage = undefined}) => previousPage;
export const getNextPageParam: GetNextPageParamFunction<Paginated<unknown>> = ({nextPage = undefined}) => nextPage;

export const usePingQuery = (options?: Omit<QueryOptions<string>, 'staleTime'>) => {
    return useQuery<string>(
        ['ping'],
        async () => {
            const {data} = await ApisauceClient.get<string>('/ping');

            return data;
        },
        {
            ...options,
            staleTime: Infinity,
        },
    );
};
