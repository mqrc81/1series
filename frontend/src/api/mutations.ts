import { ApisauceClient } from '../providers';
import { UseMutationOptions } from 'react-query/types/react/types';
import { AxiosRequestConfig } from 'axios';

export type MutationOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, TError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>

// TODO ms solve this without hacks
export const FileUploadMutation = <TData = void, TBody extends File = File>(url: string): (body: TBody) => Promise<TData> => {
    return async (body?) => {

        const {data} = await ApisauceClient.post<TData>(url, {content: await body.text()}, {timeout: undefined});

        return data;
    };
};

const postRequestConfig = <TBody>(body?: TBody): AxiosRequestConfig => {
    if (body instanceof File) {
        return {
            headers: {
                // 'content-type': 'multipart/form-data; boundary="???"',
                'content-type': body.type,
                'content-length': `${body.size}`,
            },
        };
    }
    return {};
};
