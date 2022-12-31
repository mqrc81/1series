import { useMutation } from 'react-query';
import { MutationOptions } from '../mutations';
import { FailedImdbImport } from './users.dtos';
import { ApisauceClient } from '../../providers/apisauce';

export const useImportImdbWatchlistMutation = (options?: MutationOptions<FailedImdbImport[], unknown, File>) => {
    return useMutation(
        ['users', 'importImdbWatchlist'],
        async (file: File) => {
            const {data} = await ApisauceClient.post<FailedImdbImport[]>(`/users/importImdbWatchlist`, file, {
                headers: {
                    'content-type': file.type,
                },
            });

            return data;
        },
        options,
    );
};
