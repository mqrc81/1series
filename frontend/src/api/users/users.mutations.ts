import { useMutation, useQueryClient } from 'react-query';
import { MutationOptions } from '../mutations';
import { FailedImdbImport, LoginUserDto, RegisterUserDto, TrackShowDto, User } from './users.dtos';
import { ApisauceClient } from '../../providers/apisauce';

export const useRegisterUserMutation = (options?: MutationOptions<User, unknown, RegisterUserDto>) => {
    return useMutation(
        ['users', 'register'],
        async (payload: RegisterUserDto) => {
            const {data} = await ApisauceClient.post<User>(`/users/register`, payload);

            return data;
        },
        options,
    );
};

export const useLoginUserMutation = (options?: MutationOptions<User, unknown, LoginUserDto>) => {
    return useMutation(
        ['users', 'login'],
        async (payload: LoginUserDto) => {
            const {data} = await ApisauceClient.post<User>(`/users/login`, payload);

            return data;
        },
        options,
    );
};

export const useLogoutUserMutation = (options?: MutationOptions<void>) => {
    return useMutation(
        ['users', 'logout'],
        async () => {
            await ApisauceClient.post<void>(`/users/logout`);
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
            await ApisauceClient.post<void>(`/users/trackedShows`, {showId, rating});
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
