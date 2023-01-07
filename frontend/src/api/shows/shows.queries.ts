import { QueryKey, useInfiniteQuery, useQuery } from 'react-query';
import { GenreDto, ReleaseDto, ShowDto, ShowSearchResultDto } from './shows.dtos';
import { getNextPageParam, getPreviousPageParam, InfiniteQueryOptions, QueryOptions } from '../queries';
import { Paginated } from '../dtos';
import { ApisauceClient } from '../../providers/apisauce';

export const useGetPopularShowsQuery = (options?: InfiniteQueryOptions<Paginated<{ shows: ShowDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ shows: ShowDto[] }>>(
        ['shows', 'popular'],
        async ({pageParam = 1}) => {
            const {data} = await ApisauceClient.get<Paginated<{ shows: ShowDto[] }>>(`/shows/popular`, {page: pageParam});

            return data;
        },
        {getNextPageParam, ...options},
    );
};

export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ releases: ReleaseDto[] }>>(
        ['shows', 'releases'],
        async ({pageParam = 1}) => {
            const {data} = await ApisauceClient.get<Paginated<{ releases: ReleaseDto[] }>>(`/shows/releases`, {page: pageParam});

            return data;
        },
        {getPreviousPageParam, getNextPageParam, ...options},
    );
};

export const useSearchShowsQuery = (searchTerm: string, options?: QueryOptions<ShowSearchResultDto[]> & { minParamLength: number }) => {
    return useQuery(
        ['shows', 'search', searchTerm] as QueryKey,
        async () => {
            if (searchTerm.length < options.minParamLength) return [];

            const {data} = await ApisauceClient.get<ShowSearchResultDto[]>(`/shows/search`, {searchTerm});

            return data;
        },
        options,
    );
};

export const useGetShowQuery = (id: number, options?: Omit<QueryOptions<ShowDto>, 'enabled'>) => {
    return useQuery(
        ['shows', id] as QueryKey,
        async () => {
            const {data} = await ApisauceClient.get<ShowDto>(`/shows/${id}`);

            return data;
        },
        {enabled: id > 0, ...options},
    );
};

const GENRES_TO_IGNORE = ['Talk', 'News'];
export const useGetGenresQuery = (options?: Omit<QueryOptions<GenreDto[]>, 'staleTime'>) => {
    return useQuery(
        ['shows', 'genres'] as QueryKey,
        async () => {
            const {data} = await ApisauceClient.get<GenreDto[]>(`/shows/genres`);

            return data?.filter(genre => !GENRES_TO_IGNORE.includes(genre.name));
        },
        {staleTime: Infinity, ...options},
    );
};
