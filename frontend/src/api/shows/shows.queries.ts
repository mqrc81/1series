import { useInfiniteQuery, useQuery } from 'react-query';
import { Genre, ReleaseDto, Show, ShowSearchResult } from './shows.dtos';
import { getNextPageParam, getPreviousPageParam, InfiniteQueryOptions, queryKey, QueryOptions } from '../queries';
import { Paginated } from '../dtos';
import { ApisauceClient } from '../../providers/apisauce';

export const useGetPopularShowsQuery = (options?: InfiniteQueryOptions<Paginated<{ shows: Show[] }>>) => {
    const url = `/shows/popular`;
    return useInfiniteQuery(
        queryKey(url),
        async ({pageParam = 1}) => {
            const {data} = await ApisauceClient.get<Paginated<{ shows: Show[] }>>(url, {page: pageParam});

            return data;
        },
        {getNextPageParam, ...options},
    );
};

export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    const url = `/shows/releases`;
    return useInfiniteQuery(
        queryKey(url),
        async ({pageParam = 1}) => {
            const {data} = await ApisauceClient.get<Paginated<{ releases: ReleaseDto[] }>>(url, {page: pageParam});

            return data;
        },
        {getPreviousPageParam, getNextPageParam, ...options},
    );
};

export const useSearchShowsQuery = (searchTerm: string, options?: QueryOptions<ShowSearchResult[]> & { minParamLength: number }) => {
    const url = `/shows/search`;
    return useQuery(
        queryKey(url),
        async () => {
            if (searchTerm.length < options.minParamLength) return [];

            const {data} = await ApisauceClient.get<ShowSearchResult[]>(url, {searchTerm});

            return data;
        },
        options,
    );
};

export const useGetShowQuery = (id: number, options?: Omit<QueryOptions<Show>, 'enabled'>) => {
    const url = `/shows/${id}`;
    return useQuery(
        queryKey(url),
        async () => {
            const {data} = await ApisauceClient.get<Show>(url);

            return data;
        },
        {enabled: id > 0, ...options},
    );
};

const GENRES_TO_IGNORE = ['Talk', 'News'];
export const useGetGenresQuery = (options?: Omit<QueryOptions<Genre[]>, 'staleTime'>) => {
    const url = `/shows/genres`;
    return useQuery(
        queryKey(url),
        async () => {
            const {data} = await ApisauceClient.get<Genre[]>(url);

            return data?.filter(genre => !GENRES_TO_IGNORE.includes(genre.name));
        },
        {staleTime: Infinity, ...options},
    );
};
