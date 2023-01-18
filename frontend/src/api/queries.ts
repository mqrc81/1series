import { UseQueryOptions } from 'react-query/types/react/types';
import {
    GetNextPageParamFunction,
    GetPreviousPageParamFunction,
    QueryKey,
    UseInfiniteQueryOptions,
    useQuery,
} from 'react-query';
import { ApiError, Paginated } from './dtos';
import { ApisauceClient } from '../providers/apisauce';
import { User } from './users/users.dtos';

export type QueryOptions<TQueryFnData = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseQueryOptions<TQueryFnData, ApiError, TData, TQueryKey>, 'queryKey' | 'queryFn'>
export type InfiniteQueryOptions<TQueryFnData = unknown, TData = TQueryFnData, TQueryKey extends QueryKey = QueryKey> = Omit<UseInfiniteQueryOptions<TQueryFnData, ApiError, TData, TQueryFnData, TQueryKey>, 'queryKey'>

export const getPreviousPageParam: GetPreviousPageParamFunction<Paginated<unknown>> = ({previousPage = undefined}) => previousPage;
export const getNextPageParam: GetNextPageParamFunction<Paginated<unknown>> = ({nextPage = undefined}) => nextPage;

export const useInitQuery = (options?: Omit<QueryOptions<User>, 'staleTime'>) => {
    const url = `/init`;
    return useQuery(
        queryKey(url),
        async () => {
            const {data} = await ApisauceClient.get<User>(url);

            return data;
        },
        {
            ...options,
            staleTime: Infinity,
        },
    );
};

export const queryKey = (url: string): QueryKey => url.split('/').slice(1);
