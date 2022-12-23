import { useQuery } from 'react-query';
import { GetQuery, QueryOptions } from '../index';
import { ShowDto } from './dtos';

export const useGetPopularShowsQuery = (page = 1, options?: QueryOptions<ShowDto[]>) => {
    return useQuery<ShowDto[]>(
        ['popular-shows', page],
        GetQuery(`/shows/popular`, {page}),
        options,
    );
};
