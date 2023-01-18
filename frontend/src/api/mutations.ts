import { UseMutationOptions } from 'react-query/types/react/types';
import { ApiError } from './dtos';

export type MutationOptions<TData = unknown, TVariables = void, TContext = unknown> = Omit<UseMutationOptions<TData, ApiError, TVariables, TContext>, 'mutationKey' | 'mutationFn'>
