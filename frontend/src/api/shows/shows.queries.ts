import { useInfiniteQuery } from 'react-query';
import { ReleaseDto, ShowDto } from './shows.dtos';
import { GetInfiniteQuery, InfiniteQueryOptions } from '../queries';
import { Paginated } from '../dtos';

export const useGetPopularShowsQuery = (options?: InfiniteQueryOptions<Paginated<{ shows: ShowDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ shows: ShowDto[] }>>(
        ['popular-shows'],
        GetInfiniteQuery(`/shows/popular`),
        {
            getNextPageParam: ({nextPage}) => nextPage,
            ...options,
        },
    );
};

const RELEASES_PER_REQUEST = 20;
export const useGetUpcomingReleasesQuery = (options?: InfiniteQueryOptions<Paginated<{ releases: ReleaseDto[] }>>) => {
    return useInfiniteQuery<Paginated<{ releases: ReleaseDto[] }>>(
        ['upcoming-releases'],
        GetInfiniteQuery(`/shows/releases`),
        {
            getPreviousPageParam: ({releases, previousPage}) => releases.length >= RELEASES_PER_REQUEST ? previousPage : undefined,
            getNextPageParam: ({releases, nextPage}) => releases.length >= RELEASES_PER_REQUEST ? nextPage : undefined,
            ...options,
        },
    );
};
