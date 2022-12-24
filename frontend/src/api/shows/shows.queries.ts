import { useInfiniteQuery } from 'react-query';
import { ShowDto } from './shows.dtos';
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
