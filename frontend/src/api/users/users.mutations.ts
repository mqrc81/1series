import { useMutation } from 'react-query';
import { MutationOptions, PostMutation } from '../mutations';
import { FailedImdbImport } from './users.dtos';

export const useImportImdbWatchlistMutation = (options?: MutationOptions<FailedImdbImport[], unknown, File>) => {
    return useMutation<FailedImdbImport[], unknown, File>(
        ['import-imdb-watchlist'],
        PostMutation(`/users/importImdbWatchlist`),
        {
            ...options,
        },
    );
};
