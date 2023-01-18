import { ApiErrorResponse } from 'apisauce';

export type ApiError = ApiErrorResponse<ClientError | InternalError>

export type ClientError = {
    errorMessage?: string;
}

export type InternalError = string

export type Paginated<T> = T & {
    nextPage?: number;
    previousPage?: number;
}
