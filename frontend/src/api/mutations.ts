import { UseMutationOptions } from 'react-query/types/react/types';
import { ApisauceClient } from '../providers/apisauce';

export type MutationOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, TError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>

export const FileUploadMutation = <TData = void, TBody extends File = File>(url: string): (body: TBody) => Promise<TData> => {
    return async (file?) => {

        const {data} = await ApisauceClient.post<TData>(url, file, {
            headers: {
                'content-type': file.type,
            },
        });

        return data;
    };
};
