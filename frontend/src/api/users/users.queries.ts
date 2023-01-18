import { queryKey, QueryOptions } from '../queries';
import { Show } from '../shows/shows.dtos';
import { useQuery } from 'react-query';
import { ApisauceClient } from '../../providers/apisauce';

export const useGetTrackedShowsQuery = (options: QueryOptions<Show[]> = {}) => {
    const url = `/users/trackedShows`;
    return useQuery(
        queryKey(url),
        async () => {
            const {data} = await ApisauceClient.get<Show[]>(url);

            return data;
        },
        options,
    );
};
