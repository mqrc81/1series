import { useMutation, useQueryClient } from 'react-query';
import { MutationOptions } from '../mutations';
import {
    FailedImdbImport,
    ForgotPasswordDto,
    ResetPasswordDto,
    SignUserInDto,
    SignUserUpDto,
    TrackShowDto,
    User,
} from './users.dtos';
import { ApisauceClient } from '../../providers/apisauce';
import { queryKey } from '../queries';

export const useSignUserUpMutation = (options?: MutationOptions<User, SignUserUpDto>) => {
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

export const useSignUserInMutation = (options?: MutationOptions<User, SignUserInDto>) => {
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

export const useForgotPasswordMutation = (options?: MutationOptions<void, ForgotPasswordDto>) => {
    const url = `/users/forgotPassword`;
    return useMutation(
        queryKey(url),
        async (payload: ForgotPasswordDto) => {
            await ApisauceClient.post<void>(url, payload);
        },
        options,
    );
};

export const useResetPasswordMutation = (options?: MutationOptions<void, ResetPasswordDto>) => {
    const url = `/users/resetPassword`;
    return useMutation(
        queryKey(url),
        async ({password, token}: ResetPasswordDto & { token: string }) => {
            await ApisauceClient.post<void>(url, {password}, {params: {token}});
        },
        options,
    );
};

export const useCreateTrackedShowMutation = (showId: number, options: MutationOptions<void, number> = {}) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    const url = `/users/trackedShows`;
    return useMutation(
        queryKey(url),
        async (rating: number) => {
            await ApisauceClient.post<void>(url, {showId, rating});
        },
        options,
    );
};

export const useUpdateTrackedShowMutation = (showId: number, options: MutationOptions<void, number> = {}) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    const url = `/users/trackedShows/${showId}`;
    return useMutation(
        queryKey(url),
        async (rating: number) => {
            await ApisauceClient.put<void>(url, {rating});
        },
        options,
    );
};

export const useDeleteTrackedShowMutation = (showId: number, options: MutationOptions<void, TrackShowDto> = {}) => {
    const queryClient = useQueryClient();
    options.onSuccess = (data, variables, context) => {
        void queryClient.invalidateQueries(['users', 'trackedShows']);
        options.onSuccess?.(data, variables, context);
    };

    const url = `/users/trackedShows/${showId}`;
    return useMutation(
        queryKey(url),
        async () => {
            await ApisauceClient.delete<void>(url);
        },
        options,
    );
};

export const useImportImdbWatchlistMutation = (options: MutationOptions<FailedImdbImport[], File> = {}) => {
        const queryClient = useQueryClient();
        options.onSuccess = (data, variables, context) => {
            void queryClient.invalidateQueries(['users', 'trackedShows']);
            options.onSuccess?.(data, variables, context);
        };

        const url = `/users/importImdbWatchlist`;
        return useMutation(
            queryKey(url),
            async (file: File) => {
                const {data} = await ApisauceClient.post<FailedImdbImport[]>(url, file, {headers: {'content-type': file.type}});

                return data;
            },
            options,
        );
    }
;
