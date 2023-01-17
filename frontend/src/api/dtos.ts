export type ClientError = {
    errorMessage?: string;
}

export type Paginated<T> = T & {
    nextPage?: number;
    previousPage?: number;
}
