import { useMutation, useQueryClient } from 'react-query';
import { MutationOptions } from '../mutations';
import { FailedImdbImport, SignUserInDto, SignUserUpDto, TrackShowDto, User } from './users.dtos';
import { ApisauceClient } from '../../providers/apisauce';
import { queryKey } from '../queries';

export const useSignUserUpMutation = (options?: MutationOptions<User, unknown, SignUserUpDto>) => {
    const url = `/users/signUp`;
    return useMutation(
        queryKey(url),
        async (payload: SignUserUpDto) => {
            const {data} = await ApisauceClient.post<User>(url, payload);

            return data;
        },
        options,
    );
};

export const useSignUserInMutation = (options?: MutationOptions<User, unknown, SignUserInDto>) => {
    const url = `/users/signIn`;
    return useMutation(
        queryKey(url),
        async (payload: SignUserInDto) => {
            const {data} = await ApisauceClient.post<User>(url, payload);

            return data;
        },
        options,
    );
};

export const useSignUserOutMutation = (options?: MutationOptions<void>) => {
    const url = `/users/signOut`;
    return useMutation(
        queryKey(url),
        async () => {
            await ApisauceClient.post<void>(url);
        },
        options,
    );
};

export const useCreateTrackedShowMutation = (showId: number, options?: MutationOptions<void, unknown, number>) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    return useMutation(
        ['users', 'trackedShows', showId],
        async (rating: number) => {
            const url = `/users/trackedShows`;
            await ApisauceClient.post<void>(url, {showId, rating});
        },
        options,
    );
};

export const useUpdateTrackedShowMutation = (showId: number, options?: MutationOptions<void, unknown, number>) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    return useMutation(
        ['users', 'trackedShows', showId],
        async (rating: number) => {
            await ApisauceClient.put<void>(`/users/trackedShows/${showId}`, {rating});
        },
        options,
    );
};

export const useDeleteTrackedShowMutation = (showId: number, options?: MutationOptions<void, unknown, TrackShowDto>) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    return useMutation(
        ['users', 'trackedShows', showId],
        async () => {
            await ApisauceClient.delete<void>(`/users/trackedShows/${showId}`);
        },
        options,
    );
};

export const useImportImdbWatchlistMutation = (options?: MutationOptions<FailedImdbImport[], unknown, File>) => {
        const queryClient = useQueryClient();
        options.onSuccess = (data, variables, context) => {
            void queryClient.invalidateQueries(['users', 'trackedShows']);
            options.onSuccess?.(data, variables, context);
        };

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
    }
;
