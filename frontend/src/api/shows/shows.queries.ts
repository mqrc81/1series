import { useInfiniteQuery } from 'react-query';
import { ReleaseDto, ShowDto } from './shows.dtos';
import { GetInfiniteQuery, getNextPageParam, getPreviousPageParam, InfiniteQueryOptions } from '../queries';
import { Paginated } from '../dtos';

export const useGetPopularShowsQuery = (options?: InfiniteQueryOptions<Paginated<{ shows: ShowDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ shows: ShowDto[] }>>(
        ['shows', 'popular'],
        GetInfiniteQuery(`/shows/popular`),
        {
            getNextPageParam,
            ...options,
        },
    );
};

export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ releases: ReleaseDto[] }>>(
        ['shows', 'releases'],
        GetInfiniteQuery(`/shows/releases`),
        {
            getPreviousPageParam,
            getNextPageParam,
            ...options,
        },
    );
};
