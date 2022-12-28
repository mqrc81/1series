import { ApiClient } from '../providers';
import { UseMutationOptions } from 'react-query/types/react/types';
import { AxiosRequestConfig } from 'axios';

export type MutationOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, TError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>

export const PostMutation = <TData = void, TBody = {}>(url: string, params = {}): (body: TBody) => Promise<TData> => {
    return async (body?) => {

        const {data} = await ApiClient.post<TData>(url, await postRequestBody(body));

        return data;
    };
};

const postRequestBody = async <TBody>(body?: TBody): Promise<any> => {
    if (body instanceof File) {
        return {content: await body.text()};
    }
    return body;
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
