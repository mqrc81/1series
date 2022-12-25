import { useInfiniteQuery } from 'react-query';
import { ReleaseDto, ShowDto } from './shows.dtos';
import { GetInfiniteQuery, getNextPageParam, getPreviousPageParam, InfiniteQueryOptions } from '../queries';
import { Paginated } from '../dtos';

export const useGetPopularShowsQuery = (options?: InfiniteQueryOptions<Paginated<{ shows: ShowDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ shows: ShowDto[] }>>(
        ['popular-shows'],
        GetInfiniteQuery(`/shows/popular`),
        {
            getNextPageParam,
            ...options,
        },
    );
};

export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ releases: ReleaseDto[] }>>(
        ['upcoming-releases'],
        GetInfiniteQuery(`/shows/releases`),
        {
            getPreviousPageParam,
            getNextPageParam,
            ...options,
        },
    );
};