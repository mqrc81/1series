import { UseMutationOptions } from 'react-query/types/react/types';

export type MutationOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, TError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>
