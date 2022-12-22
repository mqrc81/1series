import { useQuery } from 'react-query';
import { Apisauce } from '../index';
import { ShowDto } from './dtos';

export const useGetPopularShowsQuery = (page = 1) => {
    return useQuery(
        ['popular-shows', page],
        async () => Apisauce.get<ShowDto[]>('/shows/popular', {page}).then(({data}) => data),
    );
};
