import { useInfiniteQuery, useQuery } from 'react-query';
import { ReleaseDto, ShowDto, ShowSearchResultDto } from './shows.dtos';
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
        {
            getNextPageParam,
            ...options,
        },
    );
};

export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ releases: ReleaseDto[] }>>(
        ['shows', 'releases'],
        async ({pageParam = 1}) => {
            const {data} = await ApisauceClient.get<Paginated<{ releases: ReleaseDto[] }>>(`/shows/releases`, {page: pageParam});

            return data;
        },
        {
            getPreviousPageParam,
            getNextPageParam,
            ...options,
        },
    );
};

export const useSearchShowsQuery = (searchTerm: string, options?: QueryOptions<ShowSearchResultDto[]> & { minParamLength: number }) => {
    return useQuery<ShowSearchResultDto[]>(
        ['shows', 'search', searchTerm],
        async () => {
            if (searchTerm.length < options.minParamLength) return [];

            const {data} = await ApisauceClient.get<ShowSearchResultDto[]>(`/shows/search`, {searchTerm});

            return data;
        },
        options,
    );
};
