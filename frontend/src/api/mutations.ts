import { ApisauceClient } from '../providers';
import { UseMutationOptions } from 'react-query/types/react/types';

export type MutationOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, TError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>

// TODO ms solve this without hacks
export const FileUploadMutation = <TData = void, TBody extends File = File>(url: string): (body: TBody) => Promise<TData> => {
    return async (file?) => {

        const {data} = await ApisauceClient.post<TData>(url, file, {
            headers: {
                // 'content-type': 'multipart/form-data',
                'content-type': file.type,
                'content-length': `${file.size}`,
            },
        });

        return data;
    };
};
