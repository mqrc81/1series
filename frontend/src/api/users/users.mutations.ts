import { useMutation } from 'react-query';
import { FileUploadMutation, MutationOptions } from '../mutations';
import { FailedImdbImport } from './users.dtos';

export const useImportImdbWatchlistMutation = (options?: MutationOptions<FailedImdbImport[], unknown, File>) => {
    return useMutation<FailedImdbImport[], unknown, File>(
        ['shows', 'importImdbWatchlist'],
        FileUploadMutation(`/users/importImdbWatchlist`),
        options,
    );
};
